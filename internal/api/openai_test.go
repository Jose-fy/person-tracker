package openai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockOpenAIClient struct {
	MockResponse OpenAIResponse
	MockError    error
}

func (m *MockOpenAIClient) SendMessageGPT(message string) (OpenAIResponse, error) {
	return m.MockResponse, m.MockError
}

func TestOpenAIConnector(t *testing.T) {
	mockClient := MockOpenAIClient{
		MockResponse: OpenAIResponse{
			ID:      "chatcmpl-8RrXwXB28UGKENenrJMpqoB2nTLaQ",
			Object:  "chat.completion",
			Created: 1701651380,
			Model:   "gpt-3.5-turbo-0613",
			Choices: []struct {
				Index   int `json:"index"`
				Message struct {
					Role    string `json:"role"`
					Content string `json:"content"`
				} `json:"message"`
				FinishReason string `json:"finish_reason"`
			}{
				{
					Index: 0,
					Message: struct {
						Role    string `json:"role"`
						Content string `json:"content"`
					}{
						Role:    "assistant",
						Content: "This is a test.",
					},
					FinishReason: "stop",
				},
			},
			Usage: struct {
				PromptTokens     int `json:"prompt_tokens"`
				CompletionTokens int `json:"completion_tokens"`
				TotalTokens      int `json:"total_tokens"`
			}{
				PromptTokens:     15,
				CompletionTokens: 5,
				TotalTokens:      20,
			},
			SystemFingerprint: nil,
		},
	}

	message := "Say this is a test."
	result, _ := mockClient.SendMessageGPT(message)

	message_result := result.Choices[0].Message.Content
	assert.Equal(t, message_result, "This is a test.", "Should say this is a test")

}

func TestOpenAIResponseParser(t *testing.T){
	MockResponse :=  OpenAIResponse{
		ID:      "chatcmpl-8RrXwXB28UGKENenrJMpqoB2nTLaQ",
		Object:  "chat.completion",
		Created: 1701651380,
		Model:   "gpt-3.5-turbo-0613",
		Choices: []struct {
			Index   int `json:"index"`
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
			FinishReason string `json:"finish_reason"`
		}{
			{
				Index: 0,
				Message: struct {
					Role    string `json:"role"`
					Content string `json:"content"`
				}{
					Role:    "assistant",
					Content: "This is a test.",
				},
				FinishReason: "stop",
			},
		},
		Usage: struct {
			PromptTokens     int `json:"prompt_tokens"`
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		}{
			PromptTokens:     15,
			CompletionTokens: 5,
			TotalTokens:      20,
		},
		SystemFingerprint: nil,
	}

	ParsedOpenAIResponse := MockResponse.ParseOpenAIResponse()

	assert.Equal(t, ParsedOpenAIResponse, "This is a test.", "Should say this is a test")


}