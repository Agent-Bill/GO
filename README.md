# AgentBill Go SDK

[![CI](https://github.com/Agent-Bill/Go/workflows/CI/badge.svg)](https://github.com/Agent-Bill/Go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/Agent-Bill/Go)](https://goreportcard.com/report/github.com/Agent-Bill/Go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

OpenTelemetry-based SDK for tracking AI agent usage and billing with zero-config instrumentation.

## 📦 Installation

```bash
go get github.com/Agent-Bill/Go
```

## 📁 Project Structure

```
sdks/go/
├── .github/
│   ├── workflows/
│   │   ├── ci.yml           # Continuous integration
│   │   ├── publish.yml      # Package publishing
│   │   └── release.yml      # Release automation
│   └── ISSUE_TEMPLATE/      # Issue templates
├── examples/
│   ├── openai_basic.go      # OpenAI integration example
│   └── anthropic_basic.go   # Anthropic integration example
├── agentbill.go            # Main SDK implementation
├── tracer.go               # OpenTelemetry tracer
├── agentbill_test.go       # Main tests
├── tracer_test.go          # Tracer tests
├── go.mod                  # Go module definition
├── CHANGELOG.md            # Version history
├── CONTRIBUTING.md         # Contribution guidelines
├── SECURITY.md             # Security policy
└── README.md               # This file
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    
    "github.com/Agent-Bill/Go"
)

func main() {
    // Initialize AgentBill
    client := agentbill.Init(agentbill.Config{
        APIKey:     "your-api-key",
        CustomerID: "customer-123",
        Debug:      true,
    })

    // Wrap your OpenAI client
    openai := client.WrapOpenAI()

    // Use normally - all calls are automatically tracked!
    ctx := context.Background()
    response, err := openai.ChatCompletion(ctx, "gpt-4", []map[string]string{
        {"role": "user", "content": "Hello!"},
    })
    
    if err != nil {
        panic(err)
    }

    fmt.Printf("Response: %+v\n", response)

    // Flush telemetry
    client.Flush(ctx)
}
```

## ✨ Features

- 🚀 **Zero-config instrumentation** - Wrap your AI clients and start tracking automatically
- 📊 **Automatic tracking** - Token usage, costs, latency, and errors captured automatically
- 🔧 **Multi-provider support** - Works with OpenAI, Anthropic, and more
- 📈 **Custom signals** - Track business events with custom revenue metrics
- 🔍 **OpenTelemetry-based** - Built on industry-standard observability
- 🛡️ **Thread-safe** - Safe for concurrent use
- 🐛 **Debug mode** - Detailed logging for development

## Configuration

```go
config := agentbill.Config{
    APIKey:     "your-api-key",   // Required
    BaseURL:    "https://...",     // Optional
    CustomerID: "customer-123",    // Optional
    Debug:      true,              // Optional
}

client := agentbill.Init(config)
```

## 🧪 Development

### Running Tests
```bash
go test -v ./...
```

### Code Coverage
```bash
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
go tool cover -html=coverage.txt
```

### Code Quality
```bash
go vet ./...
go fmt ./...
```

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔒 Security

See [SECURITY.md](SECURITY.md) for security policies and reporting vulnerabilities.

## 📖 Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and release notes.

## License

MIT
