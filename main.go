package main

import (
	"main/config"
	"main/src/radarr"
	"main/src/sonarr"
	"main/types"
	"main/utils"
)

func main() {

	env, err := config.LoadConfig()
	if err != nil {
		utils.GenerateLogs("Error loading config")
		return
	}

	// Verify that Telegram credentials are available
	// Exit if either Bot token or Chat ID is missing
	if env.Telegram.Bot == "" || env.Telegram.ChatId == "" {
		utils.GenerateLogs("No telegram data available")
		return
	}

	// Initialize Telegram bot configuration
	// Uses environment variables with fallback default values
	telegramObj := types.TelegramRequest{
		Bot:    env.Telegram.Bot,
		ChatId: env.Telegram.ChatId,
	}

	// Initialize Sonarr configuration for TV show monitoring
	// Uses environment variables with fallback default values
	sonarrObj := types.Sonarr{
		Url:    env.Sonarr.Url,
		ApiKey: env.Sonarr.ApiKey,
	}

	// Check if Sonarr API key is configured
	if sonarrObj.ApiKey == "" {
		utils.GenerateLogs("No APIKEY for sonarr")
	} else {
		// Fetch TV show releases from Sonarr and send notifications via Telegram
		go sonarr.GetAllReleases(sonarrObj, telegramObj)
	}

	// Initialize Radarr configuration for movie monitoring
	// Uses environment variables with fallback default values
	radarrObj := types.Radarr{
		Url:    env.Radarr.Url,
		ApiKey: env.Radarr.ApiKey,
	}

	// Check if Radarr API key is configured
	if radarrObj.ApiKey == "" {
		utils.GenerateLogs("No APIKEY for radarr")
	} else {
		// Fetch movie releases from Radarr and send notifications via Telegram
		go radarr.GetAllReleases(radarrObj, telegramObj)
	}
}
