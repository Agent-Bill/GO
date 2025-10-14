# AgentBill Go SDK

OpenTelemetry-based SDK for automatically tracking and billing AI agent usage.

## Installation

```bash
go get github.com/agentbill/agentbill-go
```

## Quick Start

```go
package main

import (
    "context"
    "github.com/agentbill/agentbill-go"
)

func main() {
    client := agentbill.Init(agentbill.Config{
        APIKey:     "your-api-key",
        CustomerID: "customer-123",
        Debug:      true,
    })

    openai := client.WrapOpenAI()
    ctx := context.Background()
    response, err := openai.ChatCompletion(ctx, "gpt-4", []map[string]string{
        {"role": "user", "content": "Hello!"},
    })
    
    if err != nil {
        panic(err)
    }

    client.Flush(ctx)
}
```

## License

MIT