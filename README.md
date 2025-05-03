# Golendar

Golendar is a Go-based notification system that integrates with Sonarr and Radarr to send updates about new TV shows episodes and movie releases via Telegram.

## Features

- Monitors TV show releases from Sonarr
- Tracks movie releases from Radarr
- Sends notifications through Telegram with:
  - TV show details including season and episode information
  - Movie information with IMDB links
  - Show artwork/photos (when available)
  - Release overviews and descriptions

## Prerequisites

- Go 1.24.2 or higher
- A Telegram bot and chat ID
- Sonarr and/or Radarr installation with API access

## Configuration

The application uses environment variables for configuration:

### Telegram Configuration
- `TELEGRAM_BOT` - Your Telegram bot token
- `TELEGRAM_CHAT_ID` - The chat ID where notifications will be sent

### Sonarr Configuration
- `SONARR_URL` - URL of your Sonarr instance (default: http://localhost:8989)
- `SONARR_API_KEY` - Your Sonarr API key

### Radarr Configuration
- `RADARR_URL` - URL of your Radarr instance (default: http://localhost:7878)
- `RADARR_API_KEY` - Your Radarr API key

## Project Structure

```
golendar/
├── config/
│   └── config.go         # Configuration loading and management
├── constants/
│   └── constants.go      # Global constants and time formats
├── src/
│   ├── radarr/          # Radarr API integration
│   ├── sonarr/          # Sonarr API integration
│   └── telegram/        # Telegram bot messaging
├── types/
│   └── types.go         # Type definitions and structures
└── utils/
    └── utils.go         # Utility functions
```

## License

This project is available as open source under the terms of the MIT License.