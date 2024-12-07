package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

// messageCreate handles new messages sent in channels the bot has access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	switch m.Content {
	case "ping":
		sendMessage(s, m.ChannelID, "Pong!")
	case "pong":
		sendMessage(s, m.ChannelID, "Ping!")
	}
}

// presenceUpdate handles presence updates (e.g., user status/activity changes)
func presenceUpdate(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	username := getUsername(s, p)

	for _, activity := range p.Presence.Activities {
		var activityType string
		timestamp := time.UnixMilli(activity.Timestamps.StartTimestamp).Format(time.DateTime)

		switch activity.Type {
		case discordgo.ActivityTypeGame:
			activityType = "playing"
		case discordgo.ActivityTypeListening:
			activityType = "listening to"
		default:
			return
		}

		text := fmt.Sprintf("User %s is %s %s\n", username, activityType, activity.Name)
		fmt.Printf("[%s] %s", timestamp, text)

	}
}

// sendMessage is a helper to send messages in a channel
func sendMessage(s *discordgo.Session, channelId, content string) {
	if _, err := s.ChannelMessageSend(channelId, content); err != nil {
		log.Fatal("Error sending message: ", err)
	}
}

// getUsername fetches the username from the presence update
func getUsername(s *discordgo.Session, p *discordgo.PresenceUpdate) string {
	if p.User.Username != "" {
		return p.User.Username
	}

	member, err := s.State.Member(p.GuildID, p.User.ID)
	if err != nil {
		log.Fatal("Error fetching member info:", err)
		return err.Error()
	}
	return member.User.Username
}
