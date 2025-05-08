# Golendar

Golendar is a lightweight, Go-based notification service that integrates with Sonarr and Radarr media management systems. It delivers daily updates at 8:00 AM via Telegram, keeping you informed about the series episodes and movies being released that day.

## Features

- Daily notifications at 8:00 AM
- TV show episode release notifications via Sonarr integration
- Movie release notifications via Radarr integration
- Message delivery with:
  - TV show details including season and episode information
  - Movie information with IMDB links
  - Show artwork/photos (when available)
  - Release overviews and descriptions


## Prerequisites

- Go 1.24.2 or higher
- A Telegram bot and chat ID
- Sonarr and/or Radarr installation with API access
- Docker (optional, for containerized deployment)

## Usage

### Using Docker (Recommended)
You can find the Docker image here: [golendar image](https://hub.docker.com/r/brunopoiano/golendar)

#### Option 1: Pull from Docker Hub
```bash
docker run -d --name golendar --restart unless-stopped -e TZ=America/Sao_Paulo -e TELEGRAM_BOT="your_bot_token" -e TELEGRAM_CHAT_ID="your_chat_id" -e DISCORD_URL="discord_webhook_url" -e DISCORD_USERNAME="Golendar" -e SONARR_URL="sonarr_url" -e SONARR_API_KEY="sonarr_api_key" -e RADARR_URL="radarr_url" -e RADARR_API_KEY="radarr_api_key"  docker.io/brunopoiano/golendar
```

#### Option 2: Using Docker Compose
```bash
git clone https://github.com/BrunoPoiano/golendar.git
cd golendar
# update the environment variables on docker-compose.yaml file
docker compose up -d
```

### Manual Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/BrunoPoiano/golendar.git
   cd golendar
   ```
2. Set up your environment variables

3. Test the application:
   ```bash
   go run .
   ```

4. Build the application:
   ```bash
   go build -o golendar
   ```

5. Run the application:
   ```bash
   ./golendar
   ```

## Configuration

The application uses environment variables for configuration. You can set these directly in your environment or use the docker-compose.yaml file if running with Docker.

### Sonarr Configuration
- `SONARR_URL` - URL of your Sonarr instance (default: http://localhost:8989)
- `SONARR_API_KEY` - Your Sonarr API key

### Radarr Configuration
- `RADARR_URL` - URL of your Radarr instance (default: http://localhost:7878)
- `RADARR_API_KEY` - Your Radarr API key

### Discord Configuration
- `DISCORD_URL` - Your Discord webhook URL
- `DISCORD_USERNAME` - The username for the bot in Discord (default: Golendar)

### Telegram Configuration
- `TELEGRAM_BOT` - Your Telegram bot token
- `TELEGRAM_CHAT_ID` - The chat ID where notifications will be sent
