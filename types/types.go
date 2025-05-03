package types

type HttpRequest struct {
	Url    string
	Method string
	ApiKey string
	Body   []byte
}

type Config struct {
	Telegram TelegramRequest
	Sonarr   Sonarr
	Radarr   Radarr
}

// //////////////////Sonarr

type Sonarr struct {
	Url    string
	ApiKey string
}

type SonarrCalendar struct {
	SeriesId      int    `json:"seriesId"`
	SeasonNumber  int    `json:"seasonNumber"`
	EpisodeNumber int    `json:"episodeNumber"`
	Title         string `json:"title"`
	AirDate       string `json:"airDate"`
	AirDateUtc    string `json:"airDateUtc"`
	Overview      string `json:"overview"`
	Id            int    `json:"id"`
}

type SeriesInfo struct {
	Title    string      `json:"title"`
	Pictures []ImageInfo `json:"images"`
}

type ImageInfo struct {
	CoverType string `json:"coverType"`
	RemoteUrl string `json:"remoteUrl"`
}

//////////////////// Radarr

type Radarr struct {
	Url    string
	ApiKey string
}

type RadarrCalendar struct {
	Title         string `json:"title"`
	OriginalTitle string `json:"originalTitle"`
	Overview      string `json:"overview"`
	InCinemas     string `json:"inCinemas"`
	ReleaseDate   string `json:"releaseDate"`
	Id            int    `json:"id"`
	ImdbId        string `json:"imdbId"`
}

// ///////////////// Telegram
type TelegramRequest struct {
	Bot       string
	ChatId    string
	PhotoUrl  string
	Caption   string
	ParseMode string
}
