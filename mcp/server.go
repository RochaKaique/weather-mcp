package mcp

import (
	"context"
	"fmt"

	"github.com/RochaKaique/weather-mcp/internal/weather"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type WeatherForecastArgs struct {
	Lat string `json:"lat" jsonschema:"Latitude for the weather forecast"`
	Lon string `json:"lon" jsonschema:"Longitude for the weather forecast"`
}

// NewWeatherServer creates and returns an MCP server for weather forecasts
func NewWeatherServer(weatherClient *weather.Client) *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "Weather Forecast MCP",
		Version: "1.0.0",
	}, nil)

	// Register the weather_forecast tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        "weather_forecast",
		Description: "Busca previsão do tempo da Weather.gov (USA)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args WeatherForecastArgs) (*mcp.CallToolResult, any, error) {
		data, err := weatherClient.GetForecast(args.Lat, args.Lon)
		if err != nil {
			return &mcp.CallToolResult{
				IsError: true,
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("erro na API de clima: %v", err)},
				},
			}, nil, nil
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("%v", data)},
			},
		}, nil, nil
	})

	return server
}

// Start runs the weather server over stdio transport
func Start(weatherClient *weather.Client) error {
	server := NewWeatherServer(weatherClient)
	return server.Run(context.Background(), &mcp.StdioTransport{})
}
