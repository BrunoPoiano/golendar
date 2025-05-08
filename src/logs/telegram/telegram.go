package telegram

import (
	"encoding/json"
	"fmt"
	"main/types"
	"main/utils"
)

// SendTelegramPhotoMessage sends a photo message to a Telegram chat
//
// Parameters:
//   - data: Contains the chat ID, bot token, photo URL, and caption for the message
//
// Returns:
//   - None, but prints an error message to console if JSON marshaling fails
func SendTelegramPhotoMessage(data types.TelegramRequest) {

	bodyBytes, err := json.Marshal(map[string]string{
		"chat_id":    data.ChatId,
		"parse_mode": data.ParseMode,
		"photo":      data.PhotoUrl,
		"caption":    data.Caption,
	})

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", data.Bot)

	httpRequest := types.HttpRequest{
		Url:    url,
		Method: "POST",
		ApiKey: "",
		Body:   bodyBytes,
	}

	utils.HttpRequest(httpRequest)
}

// SendTelegramMessage sends a text message to a Telegram chat
//
// Parameters:
//   - data: Contains the chat ID, bot token, and caption (used as the message text)
//
// Returns:
//   - None, but prints an error message to console if JSON marshaling fails
func SendTelegramMessage(data types.TelegramRequest) {

	bodyBytes, err := json.Marshal(map[string]string{
		"chat_id":    data.ChatId,
		"parse_mode": data.ParseMode,
		"text":       data.Caption,
	})

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", data.Bot)

	httpRequest := types.HttpRequest{
		Url:    url,
		Method: "POST",
		ApiKey: "",
		Body:   bodyBytes,
	}

	utils.HttpRequest(httpRequest)
}
