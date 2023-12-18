package openai

import (
	"net/http"
	"github.com/joho/godotenv"
	"log"
	"os"
	"errors"
)

type OpenAIClient interface {
    SendMessageGPT(message string) (ChatGPTResponse, error)
    CreateEmbeddings(text string) (EmbeddingsResponse, error)
}

type RealOpenAIClient struct {
    HTTPClient *http.Client
	apiKey string
}

func getApiToken() (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return "", err
	}
	apiKey, exists := os.LookupEnv("OPENAI_API_KEY")

	if !exists {
		log.Print("No OPENAI_API_KEY found in .env file")
		return "", errors.New("No OPENAI_API_KEY found in .env file")
	}

	return apiKey, nil
}

func NewRealOpenAIClient(httpClient *http.Client) *RealOpenAIClient {
    apiKey, err := getApiToken()
    if err != nil {
        log.Fatal(err)
    }
    return &RealOpenAIClient{
        HTTPClient: httpClient,
        apiKey: apiKey,
    }
}
