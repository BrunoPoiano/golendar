package config

import (
	"main/types"
	"main/utils"
)

func LoadConfig() (*types.Config, error) {
	return &types.Config{
		Telegram: types.TelegramRequest{
			Bot:    utils.ReturnEnvVariable("TELEGRAM_BOT", "7553346015:AAGygfhQZG0Vl9D_9ky-RItP2FhGRlqFpzE"),
			ChatId: utils.ReturnEnvVariable("TELEGRAM_CHAT_ID", "1927285135"),
		},
		Sonarr: types.Sonarr{
			Url:    utils.ReturnEnvVariable("SONARR_URL", "http://192.168.3.29:8989"),
			ApiKey: utils.ReturnEnvVariable("SONARR_API_KEY", "0be88b49b6974ff1afcf4dc667c3d5fd"),
		},
		Radarr: types.Radarr{
			Url:    utils.ReturnEnvVariable("RADARR_URL", "http://localhost:7878"),
			ApiKey: utils.ReturnEnvVariable("RADARR_API_KEY", ""),
		},

		Discord: types.DiscordRequest{
			Url:      utils.ReturnEnvVariable("DISCORD_URL", ""),
			Username: utils.ReturnEnvVariable("DISCORD_USERNAME", "Golendar"),
		},
	}, nil
}
