package main

import (
	"log"

	"github.com/RochaKaique/weather-mcp/internal/weather"
	"github.com/RochaKaique/weather-mcp/mcp"
)

func main() {
	weatherClient := weather.NewClient()
	if err := mcp.Start(weatherClient); err != nil {
		log.Fatal(err)
	}
}
