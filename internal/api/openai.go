package openai

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
    "github.com/joho/godotenv"
	"os"
    "errors"
)


type OpenAIClient interface {
    SendMessageGPT(message string) (OpenAIResponse, error)
}


type RealOpenAIClient struct {
    HTTPClient *http.Client
}


func (c *RealOpenAIClient) SendMessageGPT(message string) (OpenAIResponse, error){

    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
        return OpenAIResponse{}, err
    }
    apiKey, exists := os.LookupEnv("OPENAI_API_KEY")

    if !exists {
        log.Print("No OPENAI_API_KEY found in .env file")
        return OpenAIResponse{}, errors.New("No OPENAI_API_KEY found in .env file")
    }


	request := NewOpenAIRequest(message)
	jsonData, err := json.Marshal(request)

    if err != nil {
        log.Print("Error Mashmallowing")
        return OpenAIResponse{}, err // Return an error if JSON marshaling fails
    }


	if apiKey == "" {
		log.Fatal("API key not set in environment variables")
    }

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))

	if err != nil {
        log.Print("Unable to create http request object")
		return OpenAIResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + apiKey)

	resp, err := c.HTTPClient.Do(req)

    if err != nil {
        log.Print("Unable to post http request")
        return OpenAIResponse{}, err
    }

	defer resp.Body.Close()

	var result OpenAIResponse
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&result)

	if err != nil {
        return result, err
    }

    return result, err

}


type OpenAIResponse struct {
    ID      string `json:"id"`
    Object  string `json:"object"`
    Created int    `json:"created"`
    Model   string `json:"model"`
    Choices []struct {
        Index         int `json:"index"`
        Message       struct {
            Role    string `json:"role"`
            Content string `json:"content"`
        } `json:"message"`
        FinishReason string `json:"finish_reason"`
    } `json:"choices"`
    Usage struct {
        PromptTokens     int `json:"prompt_tokens"`
        CompletionTokens int `json:"completion_tokens"`
        TotalTokens      int `json:"total_tokens"`
    } `json:"usage"`
    SystemFingerprint interface{} `json:"system_fingerprint"` // null in JSON, so using interface{}
}


func (response *OpenAIResponse) ParseOpenAIResponse() string { // Here passing by value is fine because we are not doing any changes to the response, and the response is small.
    var result string
    for _, choice := range response.Choices{
        result += choice.Message.Content
    }
    return result
}


type OpenAIRequest struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
	Temperature float32 `json:"temperature"`

}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}


func NewOpenAIRequest(message string) OpenAIRequest {
    return OpenAIRequest{
        Model: "gpt-3.5-turbo",
        Messages: []Message{
            {
                Role: "user",
                Content: message,
            },
        },
        Temperature: 0.7,
    }
}


