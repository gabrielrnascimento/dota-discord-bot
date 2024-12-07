package main

import (
	"log"

	"dota-discord-bot/internal/bot"
	"dota-discord-bot/internal/utils"
)

func main() {
	// load environment variables
	if err := utils.LoadEnv(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// start the bot
	if err := bot.Start(); err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
}
