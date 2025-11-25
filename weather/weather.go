package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type OpenMeteoCurrent struct {
	Time                string  `json:"time"`
	Interval            int     `json:"interval"`
	Temperature2m       float64 `json:"temperature_2m"`
	ApparentTemperature float64 `json:"apparent_temperature"`
	WeatherCode         int     `json:"weather_code"`
	WindSpeed10m        float64 `json:"wind_speed_10m"`
	WindDirection10m    float64 `json:"wind_direction_10m"`
	WindGusts10m        float64 `json:"wind_gusts_10m"`
}

type OpenMeteoResponse struct {
	Latitude             float64          `json:"latitude"`
	Longitude            float64          `json:"longitude"`
	Timezone             string           `json:"timezone"`
	TimezoneAbbreviation string           `json:"timezone_abbreviation"`
	Elevation            float64          `json:"elevation"`
	Current              OpenMeteoCurrent `json:"current"`
}

// GetSanFranciscoWeather queries Open-Meteo's GFS API for the current weather
// in San Francisco and returns a parsed response.
func GetSanFranciscoWeather(ctx context.Context) (*OpenMeteoResponse, error) {
	const (
		baseURL = "https://api.open-meteo.com/v1/gfs"
		lat     = 37.7749
		lon     = -122.4194
	)

	// Build query parameters
	q := url.Values{}
	q.Set("latitude", fmt.Sprintf("%.4f", lat))
	q.Set("longitude", fmt.Sprintf("%.4f", lon))

	// Ask specifically for current conditions we care about
	q.Set("current", "temperature_2m,apparent_temperature,weather_code,wind_speed_10m,wind_direction_10m,wind_gusts_10m")

	// Optional: set units + timezone
	q.Set("temperature_unit", "fahrenheit")
	q.Set("wind_speed_unit", "mph")
	q.Set("precipitation_unit", "inch")
	q.Set("timezone", "America/Los_Angeles")

	u := baseURL + "?" + q.Encode()

	// HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var data OpenMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &data, nil
}
