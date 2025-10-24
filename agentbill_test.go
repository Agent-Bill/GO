package agentbill

import (
	"context"
	"testing"
)

func TestInit(t *testing.T) {
	config := Config{
		APIKey:     "test-api-key",
		BaseURL:    "https://test.agentbill.com",
		CustomerID: "test-customer",
		Debug:      false,
	}

	client := Init(config)

	if client == nil {
		t.Fatal("Expected client to be initialized")
	}

	if client.config.APIKey != config.APIKey {
		t.Errorf("Expected APIKey %s, got %s", config.APIKey, client.config.APIKey)
	}

	if client.config.BaseURL != config.BaseURL {
		t.Errorf("Expected BaseURL %s, got %s", config.BaseURL, client.config.BaseURL)
	}

	if client.tracer == nil {
		t.Fatal("Expected tracer to be initialized")
	}
}

func TestConfigDefaults(t *testing.T) {
	config := Config{
		APIKey: "test-api-key",
	}

	client := Init(config)

	if client.config.BaseURL == "" {
		t.Error("Expected default BaseURL to be set")
	}

	if client.config.Debug {
		t.Error("Expected Debug to default to false")
	}
}

func TestTrackSignal(t *testing.T) {
	config := Config{
		APIKey:     "test-api-key",
		BaseURL:    "https://test.agentbill.com",
		CustomerID: "test-customer",
	}

	client := Init(config)

	signal := Signal{
		EventName:  "test_event",
		Revenue:    100.50,
		CustomerID: "test-customer",
		Data: map[string]interface{}{
			"key": "value",
		},
	}

	// This should not panic
	ctx := context.Background()
	err := client.TrackSignal(ctx, signal)
	if err != nil {
		t.Logf("TrackSignal returned error (expected in test environment): %v", err)
	}
}

func TestClientFlush(t *testing.T) {
	config := Config{
		APIKey:  "test-api-key",
		BaseURL: "https://test.agentbill.com",
	}

	client := Init(config)

	// Should not panic
	ctx := context.Background()
	client.Flush(ctx)
}
