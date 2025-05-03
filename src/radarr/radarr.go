package radarr

import (
	"encoding/json"
	"fmt"
	"main/constants"
	"main/src/telegram"
	"main/types"
	"main/utils"
)

// GetAllReleases retrieves all movie releases for the current day from Radarr
// and sends them as messages via Telegram.
//
// Params:
//   - radarr: Configuration details for the Radarr instance
//   - telegramBody: Template for Telegram messages with image support
//
// Returns: none
func GetAllReleases(radarr types.Radarr, telegramBody types.TelegramRequest) {

	start, end := utils.GetTimeFrame()

	httpRequest := types.HttpRequest{
		Url:    fmt.Sprintf("%s/api/v3/calendar?start=%s&end=%s&unmonitored=true", radarr.Url, start, end),
		Method: "GET",
		ApiKey: radarr.ApiKey,
		Body:   nil,
	}

	responseBody, err := utils.HttpRequest(httpRequest)
	if err != nil {
		utils.GenerateLogs(err.Error())
	}

	calendarParsed, err := parseRadarrCalendar(responseBody)
	if err != nil {
		utils.GenerateLogs(err.Error())
	}

	telegramBody.ParseMode = "markdown"

	if len(calendarParsed) > 0 {

		telegramBody.Caption = fmt.Sprintf("*Golendar* \nMovies Releasing Today:")

		telegram.SendTelegramMessage(telegramBody)

		for _, item := range calendarParsed {

			utils.GenerateLogs(fmt.Sprintf("%s | %s", item.Title, item.Overview))

			telegramBody.Caption = fmt.Sprintf("*%s* \n_%s_ \n[%s](https://www.imdb.com/title/%s/)", item.Title, item.Overview, item.Title, item.ImdbId)

			telegram.SendTelegramMessage(telegramBody)
		}

	} else {
		telegramBody.Caption = fmt.Sprintf("*Golendar* \nNo New Movies Releasing Today")

		telegram.SendTelegramMessage(telegramBody)
	}

}

// parseRadarrCalendar parses the Radarr calendar response and filters out
// movies releasing on the current day.
//
// Params:
//   - data: Raw JSON byte array from Radarr calendar API
//
// Returns:
//   - []types.RadarrCalendar: Slice of movies releasing today
//   - error: Any error that occurred during parsing
func parseRadarrCalendar(data []byte) ([]types.RadarrCalendar, error) {

	var dataParsed []types.RadarrCalendar
	var returnParsed []types.RadarrCalendar
	err := json.Unmarshal(data, &dataParsed)
	if err != nil {
		return nil, err
	}

	for _, item := range dataParsed {
		if utils.CheckSameDay(item.InCinemas, constants.UTCFormat) {
			returnParsed = append(returnParsed, item)
		}
	}

	return returnParsed, nil
}
