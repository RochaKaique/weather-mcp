package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{},
	}
}

func (c *Client) GetForecast(lat, lon string) (any, error) {
	url := fmt.Sprintf("https://api.weather.gov/points/%s,%s", lat, lon)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "go-mcp-weather")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var meta struct {
		Properties struct {
			Forecast string `json:"forecast"`
		} `json:"properties"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		return nil, err
	}

	req2, _ := http.NewRequest("GET", meta.Properties.Forecast, nil)
	req2.Header.Set("User-Agent", "go-mcp-weather")

	resp2, err := c.http.Do(req2)
	if err != nil {
		return nil, err
	}
	defer resp2.Body.Close()

	var data any
	if err := json.NewDecoder(resp2.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

