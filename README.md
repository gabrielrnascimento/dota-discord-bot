# dota-discord-bot

## Description

This is a Discord bot that provides information about Dota 2 matches and players.

## Technologies Used

- Go
- DiscordGo
- Stratz API

## Project structure
Based on the [Standard Go Project Layout](https://github.com/golang-standards/project-layout?tab=readme-ov-files)
```
dota-discord-bot/
├── cmd/
│   └── bot/
│       └── main.go          # Entry point for the bot
├── internal/
│   ├── bot/
│   │   ├── bot.go           # Bot initialization and setup
│   │   ├── handlers.go      # Event handlers (message, presence, etc.)
│   │   └── commands.go      # Command parsing and execution (if needed in the future)
│   └── utils/
│       └── config.go        # Configuration loading (e.g., environment variables)
├── .env                     # Environment variables
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
└── README.md                # Project documentation
```

## To Dos

- [ ] Post a message to a channel when a match ends with the match details
  - Match type (ranked, unranked, turbo)
  - Hero picked by the discord user

