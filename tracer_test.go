package agentbill

import (
	"context"
	"testing"
	"time"
)

func TestNewTracer(t *testing.T) {
	config := Config{
		APIKey:  "test-api-key",
		BaseURL: "https://test.agentbill.com",
		Debug:   false,
	}

	tracer := NewTracer(config)

	if tracer == nil {
		t.Fatal("Expected tracer to be initialized")
	}

	if tracer.config.APIKey != config.APIKey {
		t.Errorf("Expected APIKey %s, got %s", config.APIKey, tracer.config.APIKey)
	}

	if len(tracer.spans) != 0 {
		t.Error("Expected spans slice to be empty initially")
	}
}

func TestStartSpan(t *testing.T) {
	config := Config{
		APIKey:  "test-api-key",
		BaseURL: "https://test.agentbill.com",
	}

	tracer := NewTracer(config)

	attributes := map[string]interface{}{
		"model":    "gpt-4",
		"provider": "openai",
	}

	span := tracer.StartSpan("test.operation", attributes)

	if span == nil {
		t.Fatal("Expected span to be created")
	}

	if span.Name != "test.operation" {
		t.Errorf("Expected span name 'test.operation', got '%s'", span.Name)
	}

	if span.TraceID == "" {
		t.Error("Expected TraceID to be set")
	}

	if span.SpanID == "" {
		t.Error("Expected SpanID to be set")
	}

	if span.StartTime == 0 {
		t.Error("Expected StartTime to be set")
	}

	if model, ok := span.Attributes["model"]; !ok || model != "gpt-4" {
		t.Error("Expected attributes to be set correctly")
	}
}

func TestSpanEnd(t *testing.T) {
	config := Config{
		APIKey:  "test-api-key",
		BaseURL: "https://test.agentbill.com",
	}

	tracer := NewTracer(config)
	span := tracer.StartSpan("test.operation", nil)

	time.Sleep(10 * time.Millisecond)
	span.End()

	if span.EndTime == 0 {
		t.Error("Expected EndTime to be set after End()")
	}

	if span.EndTime <= span.StartTime {
		t.Error("Expected EndTime to be after StartTime")
	}
}

func TestSpanSetAttribute(t *testing.T) {
	config := Config{
		APIKey:  "test-api-key",
		BaseURL: "https://test.agentbill.com",
	}

	tracer := NewTracer(config)
	span := tracer.StartSpan("test.operation", nil)

	span.SetAttribute("new_key", "new_value")

	if val, ok := span.Attributes["new_key"]; !ok || val != "new_value" {
		t.Error("Expected attribute to be set")
	}
}

func TestSpanSetStatus(t *testing.T) {
	config := Config{
		APIKey:  "test-api-key",
		BaseURL: "https://test.agentbill.com",
	}

	tracer := NewTracer(config)
	span := tracer.StartSpan("test.operation", nil)

	span.SetStatus(1, "test error message")

	if code, ok := span.Status["code"].(int); !ok || code != 1 {
		t.Error("Expected status code to be 1")
	}

	if message, ok := span.Status["message"].(string); !ok || message != "test error message" {
		t.Errorf("Expected status message 'test error message', got '%s'", message)
	}
}

func TestTracerFlush(t *testing.T) {
	config := Config{
		APIKey:  "test-api-key",
		BaseURL: "https://test.agentbill.com",
	}

	tracer := NewTracer(config)
	span := tracer.StartSpan("test.operation", nil)
	span.End()

	// Should not panic (will fail to send in test, but shouldn't crash)
	ctx := context.Background()
	tracer.Flush(ctx)

	// Note: Spans are only cleared on successful flush (200 status)
	// In test environment, flush will fail, so spans won't be cleared
}
