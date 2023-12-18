package openai

import (
    "testing"
	"net/http"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func TestGetEmbeddings(t *testing.T) {



	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
		return
	}

	c := &RealOpenAIClient{
		HTTPClient: &http.Client{}, apiKey: os.Getenv("OPENAI_API_KEY"),
	}

	embedding, err := c.CreateEmbeddings("Your text string goes here", "text-embedding-ada-002")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	log.Print(embedding)

	// Add more checks here if necessary, for example to check the contents of the response
}