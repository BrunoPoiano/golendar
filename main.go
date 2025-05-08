package main

import (
	"main/config"
	"main/src/logs"
	"main/src/radarr"
	"main/src/sonarr"
	"main/types"
	"sync"
)

func main() {
	// Load application configuration from environment variables or config file
	env, err := config.LoadConfig()
	if err != nil {
		logs.MakeLog("Error loading config", nil)
		return
	}

	// Check if Telegram Bot Token and Chat ID or Discord Webhook URL are configured
	if (env.Telegram.Bot == "" || env.Telegram.ChatId == "") && (env.Discord.Url == "") {
		// If none of the notification services are configured, log a message
		logs.MakeLog("No Notification data available", nil)
	}

	// Create a WaitGroup to synchronize concurrent operations
	var wg sync.WaitGroup

	// Check if Sonarr API key is configured
	if env.Sonarr.ApiKey == "" {
		logs.MakeLog("No API_KEY for sonarr", nil)
	} else {
		// Increment the WaitGroup counter before starting the goroutine
		wg.Add(1)

		// Start Sonarr release fetching in a separate goroutine
		go func() {
			// Initialize Sonarr API client with configuration
			sonarrObj := types.Sonarr{
				Url:    env.Sonarr.Url,
				ApiKey: env.Sonarr.ApiKey,
			}

			// Decrement WaitGroup counter when this goroutine completes
			defer wg.Done()
			// Fetch and process all TV show releases from Sonarr
			sonarr.GetAllReleases(sonarrObj)
		}()
	}

	// Check if Radarr API key is configured
	if env.Radarr.ApiKey == "" {
		logs.MakeLog("No API_KEY for radarr", nil)
	} else {
		// Increment the WaitGroup counter before starting the goroutine
		wg.Add(1)
		// Start Radarr release fetching in a separate goroutine
		go func() {
			// Initialize Radarr API client with configuration
			radarrObj := types.Radarr{
				Url:    env.Radarr.Url,
				ApiKey: env.Radarr.ApiKey,
			}

			// Decrement WaitGroup counter when this goroutine completes
			defer wg.Done()
			// Fetch and process all movie releases from Radarr
			radarr.GetAllReleases(radarrObj)
		}()
	}

	// Wait for all goroutines to complete before exiting
	wg.Wait()
}
