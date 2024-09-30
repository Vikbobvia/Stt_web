package main

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-audio/wav"
	"github.com/pion/webrtc/v3"

	"github.com/asticode/go-astideepspeech"
)

const (
	httpDefaultPort   = "9000"
	defaultStunServer = "stun:stun.l.google.com:19302"
)

type TranscriptionResult struct {
	Text       string  `json:"text"`
	Confidence float32 `json:"confidence"`
	Final      bool    `json:"final"`
}

type Transcriber struct {
	model *astideepspeech.Model
	ctx   context.Context
}

func NewTranscriber(modelPath string) (*Transcriber, error) {
	model, err := astideepspeech.New(modelPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load DeepSpeech model: %v", err)
	}

	return &Transcriber{
		model: model,
		ctx:   context.Background(),
	}, nil
}

func (t *Transcriber) TranscribeAudioStream(track *webrtc.TrackRemote, dc *webrtc.DataChannel) error {
	// Create a temporary file to store audio data
	tmpFile, err := os.CreateTemp("", "deepspeech-*.wav")
	if err != nil {
		return err
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// Assuming you already have the raw PCM audio data from RTP-to-PCM conversion
	// Now write it to the WAV format file
	encoder := wav.NewEncoder(tmpFile, 16000, 16, 1, 1)
	defer encoder.Close()

	// Placeholder: Convert RTP to PCM and write to WAV format file (Implement RTP-to-PCM conversion here)

	// Transcribe the PCM data
	buffer, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return err
	}

	// Convert []byte to []int16
	var audioData []int16
	for i := 0; i < len(buffer); i += 2 {
		audioData = append(audioData, int16(binary.LittleEndian.Uint16(buffer[i:i+2])))
	}

	result, err := t.model.SpeechToText(audioData)
	if err != nil {
		return err
	}

	// Send transcription result to DataChannel
	msg, err := json.Marshal(TranscriptionResult{
		Text:       result,
		Confidence: 0.95,
		Final:      true,
	})
	if err != nil {
		return err
	}
	return dc.Send(msg)
}

type WebRTCService struct {
	stunServer  string
	transcriber *Transcriber
}

func NewWebRTCService(stunServer string, transcriber *Transcriber) *WebRTCService {
	return &WebRTCService{
		stunServer:  stunServer,
		transcriber: transcriber,
	}
}

func (w *WebRTCService) HandleAudioTrack(track *webrtc.TrackRemote, dc *webrtc.DataChannel) error {
	log.Printf("Received audio track: %s", track.ID())
	return w.transcriber.TranscribeAudioStream(track, dc)
}

func (w *WebRTCService) CreatePeerConnection() (*webrtc.PeerConnection, *webrtc.DataChannel, error) {
	pc, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{URLs: []string{w.stunServer}},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	dataChannel, err := pc.CreateDataChannel("transcription", nil)
	if err != nil {
		return nil, nil, err
	}

	pc.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		// Get the codec parameters
		codecParams := track.Codec()

		// Check if the MIME type contains "opus"
		if strings.Contains(strings.ToLower(codecParams.MimeType), "opus") {
			go func() {
				if err := w.HandleAudioTrack(track, dataChannel); err != nil {
					log.Printf("Error handling track: %v", err)
				}
			}()
		}
	})

	return pc, dataChannel, nil
}

type SessionRequest struct {
	Offer string `json:"offer"`
}

type SessionResponse struct {
	Answer string `json:"answer"`
}

func MakeSessionHandler(webrtcService *WebRTCService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req SessionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pc, dc, err := webrtcService.CreatePeerConnection()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer pc.Close()
		defer dc.Close()

		err = pc.SetRemoteDescription(webrtc.SessionDescription{
			SDP:  req.Offer,
			Type: webrtc.SDPTypeOffer,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		answer, err := pc.CreateAnswer(nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = pc.SetLocalDescription(answer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := SessionResponse{
			Answer: answer.SDP,
		}
		json.NewEncoder(w).Encode(resp)
	})
}

func main() {
	httpPort := flag.String("http.port", httpDefaultPort, "HTTP listen port")
	stunServer := flag.String("stun.server", defaultStunServer, "STUN server URL")
	modelPath := flag.String("stt.model", "./model", "Path to the Coqui STT model directory")
	flag.Parse()

	transcriber, err := NewTranscriber(*modelPath)
	if err != nil {
		log.Fatalf("Failed to create transcriber: %v", err)
	}

	webrtcService := NewWebRTCService(*stunServer, transcriber)

	http.Handle("/session", MakeSessionHandler(webrtcService))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/index.html")
	})
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./web"))))

	errors := make(chan error, 2)
	go func() {
		log.Printf("Starting server on port %s", *httpPort)
		errors <- http.ListenAndServe(fmt.Sprintf(":%s", *httpPort), nil)
	}()

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		errors <- fmt.Errorf("received signal %v", <-interrupt)
	}()

	err = <-errors
	log.Printf("Server stopped: %v", err)
}
