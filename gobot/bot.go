package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var _ = time.Now

const (
	TestChannel = "413813088018497536"
)

var conerr error = fmt.Errorf("No connection established")

type Bot struct {
	token      string
	Connection *discordgo.Session
}

func NewBot(token string) *Bot {
	b := new(Bot)
	b.token = token

	return b
}

func (b *Bot) Close() {
	b.Connection.Close()
}

func (b *Bot) Send(channel, message string) error {
	if b.Connection == nil {
		return conerr
	}

	_, err := b.Connection.ChannelMessageSend(TestChannel, message)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) Connect() error {
	se, err := discordgo.New("Bot " + b.token)
	if err != nil {
		return err
	}
	b.Connection = se
	err = se.Open()
	if err != nil {
		return err
	}
	b.init()
	return nil
}

func (b *Bot) init() {
	b.Connection.AddHandler(messageCreate)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
