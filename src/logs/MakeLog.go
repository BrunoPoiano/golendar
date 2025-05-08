package logs

import (
	"fmt"
	"main/config"
	"main/src/logs/discord"
	"main/src/logs/telegram"
	"main/types"
	"time"
)

// MakeLog generates a log message and sends notifications to Telegram and Discord
// if configured.
//
// Params:
//   - log: The log message string
//   - logObj: Pointer to a MessageType struct containing data for notifications.
//     If nil, only the simple log is written.
func MakeLog(log string, logObj *types.MessageType) {

	// simple Log
	if log != "" {
		now := time.Now()
		timeFormated := fmt.Sprintf("%02d/%02d/%d %02d:%02d:%02d", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())
		println(timeFormated, "|", log)
	}

	if logObj == nil {
		return
	}

	// Notification log
	env, err := config.LoadConfig()
	if err != nil {
		return
	}

	// Send Telegram Notification
	if env.Telegram.Bot != "" && env.Telegram.ChatId != "" {
		telegramObj := types.TelegramRequest{
			Bot:    env.Telegram.Bot,
			ChatId: env.Telegram.ChatId,
		}

		telegramObj.ParseMode = "markdown"
		if logObj.ParseMode != "" {
			telegramObj.ParseMode = logObj.ParseMode
		}

		telegramObj.Caption = logObj.Message
		if logObj.PhotoUrl != "" {
			telegramObj.PhotoUrl = logObj.PhotoUrl
			telegram.SendTelegramPhotoMessage(telegramObj)
		} else {
			telegram.SendTelegramMessage(telegramObj)
		}
	}

	// Send Discord Notification
	if env.Discord.Url != "" {
		discordObj := types.DiscordRequest{
			Url:      env.Discord.Url,
			Username: env.Discord.Username,
		}

		discordObj.Content = logObj.Message
		discordObj.EmbedTitle = logObj.PhotoCaption
		discordObj.EmbedPhotoUrl = logObj.PhotoUrl
		discord.SendDiscordMessage(discordObj)
	}

}
