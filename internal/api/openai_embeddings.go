package openai

import (
	"bytes"
	"encoding/json"
		"log"
	"net/http"
)

func (c *RealOpenAIClient) CreateEmbeddings(text string, model string) (EmbeddingsResponse, error) {

	request := NewEmbeddingsRequest(text, model)
	jsonData, err := json.Marshal(request)

	if err != nil {
		log.Print("Error Mashmallowing")
		return EmbeddingsResponse{}, err // Return an error if JSON marshaling fails
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Print("Unable to create http request object")
		return EmbeddingsResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)


	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		log.Print("Unable to post http request")
		return EmbeddingsResponse{}, err
	}

	defer resp.Body.Close()

	var result EmbeddingsResponse
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&result)


	if err != nil {
		log.Print("Unable to decode response")
		return result, err
	}

	return result, err
}

type EmbeddingsResponse struct {
	Data   []Data `json:"data"`
	Model  string `json:"model"`
	Object string `json:"object"`
	Usage  Usage  `json:"usage"`
}

type Data struct {
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
	Object    string    `json:"object"`
}

type Usage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

func (response *EmbeddingsResponse) ParseEmbeddingResponse() []float64 {
	return response.Data[0].Embedding
}

type EmbeddingsRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

func NewEmbeddingsRequest(text string, model string) EmbeddingsRequest {
	return EmbeddingsRequest{
		Input: text,
		Model: model,
	}
}
