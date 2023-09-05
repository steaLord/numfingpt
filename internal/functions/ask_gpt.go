package functions

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func AskGpt(body string, apiKey string) (string, error) {

	if apiKey == "" {
		return "", fmt.Errorf("Missing API key")
	}
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: body + " Add 4 sentences to this text giving analytics",
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v\n", err)
	}

	return fmt.Sprintf("%s\n\n%s", body, resp.Choices[0].Message.Content), nil
}
