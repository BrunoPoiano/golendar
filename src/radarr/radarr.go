package radarr

import (
	"encoding/json"
	"fmt"
	"main/constants"
	"main/src/logs"
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
func GetAllReleases(radarr types.Radarr) {

	start, end := utils.GetTimeFrame()

	httpRequest := types.HttpRequest{
		Url:    fmt.Sprintf("%s/api/v3/calendar?start=%s&end=%s&unmonitored=true", radarr.Url, start, end),
		Method: "GET",
		ApiKey: radarr.ApiKey,
		Body:   nil,
	}

	responseBody, err := utils.HttpRequest(httpRequest)
	if err != nil {
		logs.MakeLog(err.Error(), nil)
	}

	calendarParsed, err := parseRadarrCalendar(responseBody)
	if err != nil {
		logs.MakeLog(err.Error(), nil)
	}

	messageType := types.MessageType{
		Message: fmt.Sprintf("*Golendar* \nMovies Releasing Today:"),
	}

	if len(calendarParsed) > 0 {

		logs.MakeLog("Golendar | Movies Releasing Today:", &messageType)

		for _, item := range calendarParsed {

			log := fmt.Sprintf("%s | %s", item.Title, item.Overview)
			messageType.Message = fmt.Sprintf("*%s* \n_%s_ \n[%s](https://www.imdb.com/title/%s/)", item.Title, item.Overview, item.Title, item.ImdbId)
			logs.MakeLog(log, &messageType)
		}

	} else {
		log := ("Golendar | No New Movies Releasing Today")
		messageType.Message = fmt.Sprintf("*Golendar* \nNo New Movies Releasing Today")
		logs.MakeLog(log, &messageType)
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
