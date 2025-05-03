package config

import (
	"main/types"
	"main/utils"
)

func LoadConfig() (*types.Config, error) {
	return &types.Config{
		Telegram: types.TelegramRequest{
			Bot:    utils.ReturnEnvVariable("TELEGRAM_BOT", ""),
			ChatId: utils.ReturnEnvVariable("TELEGRAM_CHAT_ID", ""),
		},
		Sonarr: types.Sonarr{
			Url:    utils.ReturnEnvVariable("SONARR_URL", "http://localhost:8989"),
			ApiKey: utils.ReturnEnvVariable("SONARR_API_KEY", ""),
		},
		Radarr: types.Radarr{
			Url:    utils.ReturnEnvVariable("RADARR_URL", "http://localhost:7878"),
			ApiKey: utils.ReturnEnvVariable("RADARR_API_KEY", ""),
		},
	}, nil
}
