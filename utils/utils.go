package utils

import (
	"bytes"
	"fmt"
	"io"
	"main/constants"
	"main/types"
	"net/http"
	"os"
	"time"
)

// ReturnEnvVariable returns the value of the environment variable specified by key.
// If the environment variable is not set, it returns the default_value.
//
// Parameters:
//   - key: the name of the environment variable
//   - default_value: the value to return if the environment variable is not set
//
// Returns:
//   - string: the value of the environment variable or the default value
func ReturnEnvVariable(key, default_value string) string {
	envKey := os.Getenv(key)
	if envKey == "" {
		return default_value
	}

	return envKey
}

// HttpRequest makes an HTTP request using the provided HttpRequest data.
//
// Parameters:
//   - data: an HttpRequest struct containing request details (Method, Url, Body, ApiKey)
//
// Returns:
//   - []byte: the response body
//   - error: any error encountered during the request
func HttpRequest(data types.HttpRequest) ([]byte, error) {
	responseBody := bytes.NewBuffer(data.Body)

	request, err := http.NewRequest(data.Method, data.Url, responseBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	if data.ApiKey != "" {
		request.Header.Add("X-Api-Key", data.ApiKey)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetTimeFrame returns a start and end time for a 24-hour period.
// The start time is the current day at 03:00:00 UTC.
// The end time is the next day at 02:59:59.999 UTC.
//
// Returns:
//   - string: formatted start time (YYYY-MM-DDT03:00:00.000Z)
//   - string: formatted end time (YYYY-MM-DDT02:59:59.999Z)
func GetTimeFrame() (string, string) {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)

	start := fmt.Sprintf("%d-%02d-%02dT%s", now.Year(), now.Month(), now.Day(), constants.DefaultStartTime)
	end := fmt.Sprintf("%d-%02d-%02dT%s", now.Year(), now.Month(), tomorrow.Day(), constants.DefaultEndTime)

	return start, end
}

// CheckSameDay checks if the provided date string represents the same day as the current day.
//
// Parameters:
//   - dateStr: a date string to be parsed
//   - parser: the layout string to use for parsing the date
//
// Returns:
//   - bool: true if the date is the same as the current day, false otherwise or if parsing fails
func CheckSameDay(dateStr, parser string) bool {
	now := time.Now()
	date, err := time.Parse(parser, dateStr)
	if err != nil {
		return false
	}

	return date.Day() == now.Day()
}

// SeriesFormat formats season and episode numbers into standard TV series format (S01E02).
//
// Parameters:
//   - season: the season number
//   - epNumber: the episode number
//
// Returns:
//   - string: formatted season and episode string (e.g., "S01E02")
func SeriesFormat(season, epNumber int) string {
	return fmt.Sprintf("S%02dE%02d", season, epNumber)
}
