from flask import Flask, request, jsonify
import speech_recognition as sr
import pyaudio
import io

app = Flask(__name__)

@app.route('/speech-to-text', methods=['POST'])
def speech_to_text():
    audio_file = request.files.get('file')
    recognizer = sr.Recognizer()

    if audio_file:
        try:
            with sr.AudioFile(audio_file) as source:
                audio_data = recognizer.record(source)
                text = recognizer.recognize_google(audio_data)
            return jsonify({"transcript": text})
        except sr.UnknownValueError:
            return jsonify({"error": "Speech recognition could not understand audio"})
        except sr.RequestError as e:
            return jsonify({"error": f"Could not request results from Google Speech Recognition service; {e}"})
        except Exception as e:
            return jsonify({"error": str(e)})
    
    return jsonify({"error": "No audio data provided"}), 400

@app.route('/live-speech-to-text', methods=['POST'])
def live_speech_to_text():
    recognizer = sr.Recognizer()
    mic = sr.Microphone()

    try:
        with mic as source:
            print("Listening...")
            audio = recognizer.listen(source)
            text = recognizer.recognize_google(audio)
            print("You said: " + text)
            return jsonify({"transcript": text})
    except sr.RequestError:
        return jsonify({"error": "API unavailable"})
    except sr.UnknownValueError:
        return jsonify({"error": "Unable to recognize speech"})
    except Exception as e:
        return jsonify({"error": str(e)})

if __name__ == '__main__':
    app.run(debug=True)

