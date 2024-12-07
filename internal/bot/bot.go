package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func Start() error {
	// load the bot token from the environment variable
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		return fmt.Errorf("DISCORD_BOT_TOKEN is required")
	}

	// create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("Error creating Discord session: " + err.Error())
	}

	// register event handlers
	session.AddHandler(messageCreate)
	session.AddHandler(presenceUpdate)

	// set bot intents
	session.Identify.Intents = discordgo.IntentsAll

	// open a websocket connection to Discord and begin listening.
	err = session.Open()
	if err != nil {
		return fmt.Errorf("Error opening connection: " + err.Error())
	}

	defer func() {
		if err := session.Close(); err != nil {
			fmt.Println("Error closing session: ", err)
		}
	}()

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// Wait for termination signal to gracefully shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return nil
}
