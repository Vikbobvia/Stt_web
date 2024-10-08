package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "strings"
)

type SoundFileData struct {
    CreatorName string `json:"creator_name"`
    FileName    string `json:"file_name"`
    FileContent string `json:"file_content"`
}

// Function to count words in a string
func countingWords(content string) int {
    words := strings.Fields(content) // Split the string into words
    return len(words)                // Return the count of words
}

// Handler for counting words
func countWordsHandler(c *fiber.Ctx) error {
    var data SoundFileData

    // Parse JSON request body
    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request",
        })
    }

    // Count the words in the file content
    wordCount := countingWords(data.FileContent)

    // Return the result
    return c.JSON(fiber.Map{
        "status":       "File received successfully",
        "file_name":    data.FileName + " + go",
        "creator_name": data.CreatorName,
        "word_count":   wordCount,
    })
}

func main() {
    app := fiber.New()

    // Enable CORS middleware
    app.Use(cors.New())

    // API route to count words and return file status
    app.Post("/api/countWords", countWordsHandler)

    // Start Fiber server on port 8080
    app.Listen(":8080")
}
