package discord

import (
	"encoding/json"
	"fmt"
	"main/types"
	"main/utils"
)

func SendDiscordMessage(data types.DiscordRequest) {
	var body map[string]interface{}

	body = map[string]interface{}{
		"content":  data.Content,
		"username": data.Username,
	}

	if data.EmbedPhotoUrl != "" {
		body["embeds"] = []interface{}{
			map[string]interface{}{
				"title": data.EmbedTitle,
				"image": map[string]interface{}{
					"url": data.EmbedPhotoUrl,
				},
			},
		}
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	httpRequest := types.HttpRequest{
		Url:    data.Url,
		Method: "POST",
		Body:   bodyBytes,
	}

	utils.HttpRequest(httpRequest)
}
