# Weather Forecast MCP Server

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.25.5-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Model Context Protocol](https://img.shields.io/badge/MCP-v1.1.0-2563EB?style=flat-square)](https://modelcontextprotocol.io)
[![Build Status](https://img.shields.io/badge/Build-Passing-success?style=flat-square)](https://github.com/RochaKaique/weather-mcp)

A robust Model Context Protocol (MCP) server implementation for fetching weather forecasts using the Weather.gov API.

[Features](#features) • [Quick Start](#quick-start) • [Usage](#usage) • [Architecture](#architecture) • [Contributing](#contributing)

</div>

---

## Features

✨ **MCP Server Implementation** - Full Model Context Protocol compliance using the official Go SDK  
🌤️ **Weather Forecasts** - Real-time weather data from NOAA's Weather.gov API  
📍 **Coordinate-based Lookups** - Query weather by latitude and longitude  
🔄 **Stdio Transport** - Seamless integration with MCP clients via stdin/stdout  
⚡ **Lightweight & Fast** - Minimal dependencies, optimized for performance  
🛠️ **Type-safe** - Full type safety with Go's strict typing system  

---

## Quick Start

### Prerequisites

- Go 1.25.5 or higher
- A terminal or IDE with Go support

### Installation

```bash
# Clone the repository
git clone https://github.com/RochaKaique/weather-mcp.git
cd weather-mcp

# Download dependencies
go mod download

# Build the project
go build -o weather-mcp ./cmd
```

### Running the Server

```bash
# Run the server
./weather-mcp
```

The server will now listen on stdin/stdout and is ready to receive MCP requests from clients.

---

## Usage

### Tool: `weather_forecast`

Get weather forecast data for a specific location.

**Parameters:**
- `lat` (string, required) - Latitude coordinate for the weather forecast
- `lon` (string, required) - Longitude coordinate for the weather forecast

**Example Request:**

```json
{
  "method": "tools/call",
  "params": {
    "name": "weather_forecast",
    "arguments": {
      "lat": "40.7128",
      "lon": "-74.0060"
    }
  }
}
```

**Example Response:**

```json
{
  "content": [
    {
      "type": "text",
      "text": "Forecast data for New York City..."
    }
  ],
  "isError": false
}
```

### Connecting with an MCP Client

```go
import (
  "context"
  "os/exec"
  "github.com/modelcontextprotocol/go-sdk/mcp"
)

client := mcp.NewClient(
  &mcp.Implementation{Name: "my-client", Version: "1.0.0"},
  nil,
)

transport := &mcp.CommandTransport{
  Command: exec.Command("./weather-mcp"),
}

session, err := client.Connect(context.Background(), transport, nil)
if err != nil {
  log.Fatal(err)
}
defer session.Close()

// Call the weather forecast tool
result, err := session.CallTool(context.Background(), &mcp.CallToolParams{
  Name: "weather_forecast",
  Arguments: map[string]any{
    "lat": "40.7128",
    "lon": "-74.0060",
  },
})
```

---

## Architecture

### Project Structure

```
weather-mcp/
├── cmd/
│   └── main.go           # Entry point
├── internal/
│   └── weather/
│       └── client.go     # Weather API client
├── mcp/
│   └── server.go         # MCP server implementation
├── go.mod                # Module definition
├── go.sum                # Dependency checksums
└── README.md             # This file
```

### Core Components

#### MCP Server (`mcp/server.go`)

Implements the Model Context Protocol with:
- Tool registration using `mcp.AddTool`
- Stdio transport for seamless client communication
- Type-safe argument handling with JSON schema support
- Structured error handling

#### Weather Client (`internal/weather/client.go`)

Handles API communication with Weather.gov:
- Fetches points data from coordinate-based endpoints
- Retrieves detailed forecast information
- Error handling and resilience

---

## Technology Stack

| Technology | Version | Purpose |
|-----------|---------|---------|
| **Go** | 1.25.5 | Core programming language |
| **MCP SDK** | 1.1.0 | Model Context Protocol implementation |
| **jsonschema-go** | 0.3.0 | JSON Schema support |
| **oauth2** | 0.30.0 | OAuth2 support (dependency) |

---

## Development

### Building from Source

```bash
# Install dependencies
go mod tidy

# Build the binary
go build -o weather-mcp ./cmd

# Run tests (if available)
go test ./...
```

### Project Structure

The project follows Go best practices:
- `cmd/` - Executable entry points
- `internal/` - Private packages not exported
- `mcp/` - MCP server implementation

---

## API Reference

### Weather Forecast Tool

**Name:** `weather_forecast`

**Description:** Fetches weather forecast from Weather.gov (USA only)

**Input Schema:**
```json
{
  "type": "object",
  "properties": {
    "lat": {
      "type": "string",
      "description": "Latitude for the weather forecast"
    },
    "lon": {
      "type": "string",
      "description": "Longitude for the weather forecast"
    }
  },
  "required": ["lat", "lon"]
}
```

**Output:** Weather forecast data in text format

**Notes:**
- Requires valid USA coordinates
- Data sourced from NOAA's Weather.gov API
- Returns detailed forecast including temperature, conditions, and alerts

---

## Error Handling

The server provides clear error messages for various scenarios:

- **Invalid Coordinates:** Missing or malformed latitude/longitude
- **API Errors:** Issues communicating with Weather.gov
- **Network Issues:** Connection timeouts or failures

All errors are returned in the MCP response with `IsError: true`.

---

## Performance Considerations

- **Caching:** Consider implementing forecast caching to reduce API calls
- **Rate Limiting:** Weather.gov has rate limits; implement backoff strategies
- **Timeout:** API calls include reasonable timeout values
- **Concurrency:** Supports multiple concurrent forecast requests

---

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### How to Contribute

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Model Context Protocol](https://modelcontextprotocol.io) - Official MCP specification
- [Go MCP SDK](https://github.com/modelcontextprotocol/go-sdk) - Official Go implementation
- [NOAA Weather.gov](https://weather.gov) - Weather data provider
- Go community for excellent tooling and documentation

---

## Support

If you encounter any issues or have questions:

- 📝 Check existing [Issues](https://github.com/RochaKaique/weather-mcp/issues)
- 💬 Open a [Discussion](https://github.com/RochaKaique/weather-mcp/discussions)
- 🐛 Report bugs with detailed information
- 📖 Review the [MCP Documentation](https://modelcontextprotocol.io/docs)

---

<div align="center">

Made with ❤️ by [Kaique Rocha](https://github.com/RochaKaique)

**[⬆ back to top](#weather-forecast-mcp-server)**

</div>
