package sonarr

import (
	"encoding/json"
	"fmt"
	"main/constants"
	"main/src/logs"
	"main/types"
	"main/utils"
)

// GetAllReleases fetches TV show episodes that are releasing today from Sonarr
// and prepares Telegram message content based on the results.
//
// Params:
//   - sonarr: Configuration for the Sonarr API connection
//   - telegramBody: The Telegram message request object to be populated
func GetAllReleases(sonarr types.Sonarr) {

	start, end := utils.GetTimeFrame()

	httpRequest := types.HttpRequest{
		Url:    fmt.Sprintf("%s/api/v3/calendar?start=%s&end=%s&unmonitored=true", sonarr.Url, start, end),
		Method: "GET",
		ApiKey: sonarr.ApiKey,
		Body:   nil,
	}

	responseBody, err := utils.HttpRequest(httpRequest)
	if err != nil {
		logs.MakeLog(err.Error(), nil)
	}

	calendarParsed, err := parseSonarrCalendar(responseBody)
	if err != nil {
		logs.MakeLog(err.Error(), nil)
	}

	messageType := types.MessageType{
		Message: fmt.Sprintf("*Golendar* \nEpisodes Releasing Today:"),
	}

	if len(calendarParsed) > 0 {

		logs.MakeLog("Golendar | Episodes Releasing Today", &messageType)

		for _, item := range calendarParsed {
			seriesParsed, err := getSeriesInfo(sonarr, item)
			if err != nil {
				logs.MakeLog(err.Error(), nil)
				continue
			}

			seriesFormat := utils.SeriesFormat(item.SeasonNumber, item.EpisodeNumber)

			log := fmt.Sprintf("%s (%s) | %s - %s | %s", seriesParsed.Title, seriesParsed.AirTime, seriesFormat, item.Title, item.Overview)
			messageType.Message = fmt.Sprintf("*%s* (%sh) \n%s - %s \n\n%s", seriesParsed.Title, seriesParsed.AirTime, seriesFormat, item.Title, item.Overview)
			messageType.PhotoUrl = seriesParsed.Pictures[1].RemoteUrl
			messageType.PhotoCaption = seriesParsed.Title
			logs.MakeLog(log, &messageType)
		}

	} else {
		log := "Golendar | No New Series episodes Releasing Today"
		messageType.Message = fmt.Sprintf("*Golendar* \nNo New Series episodes Releasing Today")
		logs.MakeLog(log, &messageType)
	}
}

// getSeriesInfo retrieves detailed information about a TV series from Sonarr.
//
// Params:
//   - sonarr: Configuration for the Sonarr API connection
//   - sonarrCalendar: Calendar item containing the series ID to look up
//
// Returns:
//   - types.SeriesInfo: The parsed series information
//   - error: Any error encountered during the operation
func getSeriesInfo(sonarr types.Sonarr, sonarrCalendar types.SonarrCalendar) (types.SeriesInfo, error) {

	httpRequest := types.HttpRequest{
		Url:    fmt.Sprintf("%s/api/v3/series/%d", sonarr.Url, sonarrCalendar.SeriesId),
		Method: "GET",
		ApiKey: sonarr.ApiKey,
		Body:   nil,
	}

	resp, err := utils.HttpRequest(httpRequest)
	if err != nil {
		return types.SeriesInfo{}, err
	}

	seriesParsed, err := parseSeriesInfo(resp)
	if err != nil {
		return types.SeriesInfo{}, err
	}

	return seriesParsed, nil
}

// parseRadarrCalendar parses the calendar data from Sonarr and filters items
// that are releasing today.
//
// Params:
//   - data: Raw JSON bytes from Sonarr API response
//
// Returns:
//   - []types.SonarrCalendar: Array of calendar items releasing today
//   - error: Any error encountered during parsing
func parseSonarrCalendar(data []byte) ([]types.SonarrCalendar, error) {

	var dataParsed []types.SonarrCalendar
	var returnParsed []types.SonarrCalendar
	err := json.Unmarshal(data, &dataParsed)
	if err != nil {
		return nil, err
	}

	for _, item := range dataParsed {
		if utils.CheckSameDay(item.AirDate, constants.TimeFormat) {
			returnParsed = append(returnParsed, item)
		}
	}

	return returnParsed, nil
}

// parseSeriesInfo unmarshals series information from JSON response.
//
// Params:
//   - data: Raw JSON bytes from Sonarr API response containing series info
//
// Returns:
//   - types.SeriesInfo: The parsed series information
//   - error: Any error encountered during parsing
func parseSeriesInfo(data []byte) (types.SeriesInfo, error) {

	var dataParsed types.SeriesInfo

	err := json.Unmarshal(data, &dataParsed)
	if err != nil {
		return types.SeriesInfo{}, err
	}

	return dataParsed, nil
}
