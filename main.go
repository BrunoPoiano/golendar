package main

import (
	"main/config"
	"main/src/radarr"
	"main/src/sonarr"
	"main/types"
	"main/utils"
	"sync"
)

func main() {
	// Load application configuration from environment variables or config file
	env, err := config.LoadConfig()
	if err != nil {
		utils.GenerateLogs("Error loading config")
		return
	}

	// Verify that Telegram credentials are available
	if env.Telegram.Bot == "" || env.Telegram.ChatId == "" {
		utils.GenerateLogs("No telegram data available")
		return
	}

	// Initialize Telegram notification object with credentials
	telegramObj := types.TelegramRequest{
		Bot:    env.Telegram.Bot,
		ChatId: env.Telegram.ChatId,
	}

	// Create a WaitGroup to synchronize concurrent operations
	var wg sync.WaitGroup

	// Check if Sonarr API key is configured
	if env.Sonarr.ApiKey == "" {
		utils.GenerateLogs("No API_KEY for sonarr")
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
			sonarr.GetAllReleases(sonarrObj, telegramObj)
		}()
	}

	// Check if Radarr API key is configured
	if env.Radarr.ApiKey == "" {
		utils.GenerateLogs("No API_KEY for radarr")
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
			radarr.GetAllReleases(radarrObj, telegramObj)
		}()
	}

	// Wait for all goroutines to complete before exiting
	wg.Wait()
}
