package types

type HttpRequest struct {
	Url    string
	Method string
	ApiKey string
	Body   []byte
}

type Config struct {
	Telegram TelegramRequest
	Discord  DiscordRequest
	Sonarr   Sonarr
	Radarr   Radarr
}

type MessageType struct {
	Message      string
	PhotoUrl     string
	PhotoCaption string
	ParseMode    string
}

// ///////////////// Telegram
type TelegramRequest struct {
	Bot       string
	ChatId    string
	PhotoUrl  string
	Caption   string
	ParseMode string
}

// ///////////// Discord
type DiscordRequest struct {
	Url           string
	Content       string
	Username      string
	EmbedTitle    string
	EmbedPhotoUrl string
}

////////////////////Sonarr

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
	AirTime  string      `json:"airTime"`
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
