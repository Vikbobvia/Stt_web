package main

import (
    "bytes"
    "fmt"
    "io"
    "log"
    "mime/multipart"
    "net/http"
    "os"

    "github.com/gofiber/fiber/v2"
)

func uploadFile(url string, filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    var b bytes.Buffer
    writer := multipart.NewWriter(&b)
    part, err := writer.CreateFormFile("file", filePath)
    if err != nil {
        return "", fmt.Errorf("error creating form file: %v", err)
    }

    _, err = io.Copy(part, file)
    if err != nil {
        return "", fmt.Errorf("error copying file: %v", err)
    }
    writer.Close()

    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
        return "", fmt.Errorf("error creating request: %v", err)
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("error making request: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response: %v", err)
    }

    return string(body), nil
}

func main() {
    app := fiber.New()

    app.Static("/", "./public")

    app.Post("/speech-to-text", func(c *fiber.Ctx) error {
        file, err := c.FormFile("file")
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "file not found"})
        }

        filePath := fmt.Sprintf("./%s", file.Filename)
        if err := c.SaveFile(file, filePath); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save file"})
        }

        pythonAPIUrl := "http://localhost:5000/speech-to-text"
        transcript, err := uploadFile(pythonAPIUrl, filePath)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }

        return c.JSON(fiber.Map{"transcript": transcript})
    })

    log.Fatal(app.Listen(":3000"))
}

