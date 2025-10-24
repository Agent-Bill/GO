package main

import (
	"context"
	"fmt"
	"log"
	"os"

	agentbill "github.com/Agent-Bill/Go"
	anthropic "github.com/anthropics/anthropic-sdk-go"
)

func main() {
	// Initialize AgentBill
	client := agentbill.Init(agentbill.Config{
		APIKey:     os.Getenv("AGENTBILL_API_KEY"),
		BaseURL:    "https://agentbill.com",
		CustomerID: "customer-123",
		Debug:      true,
	})

	// Create Anthropic client
	anthropicClient := anthropic.NewClient(
		anthropic.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
	)

	// Wrap Anthropic client with AgentBill tracking
	wrappedClient := client.WrapAnthropic(anthropicClient)

	// Make a request - automatically tracked
	ctx := context.Background()
	response, err := wrappedClient.Messages.Create(ctx, anthropic.MessageCreateParams{
		Model: anthropic.ModelClaude3_5Sonnet20241022,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock("What is the capital of France?")),
		},
		MaxTokens: 1024,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %s\n", response.Content[0].Text)

	// Flush any pending traces
	client.Flush(ctx)
}
