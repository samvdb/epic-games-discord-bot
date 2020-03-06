package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"time"
)

type DiscordClient struct {
	session    *discordgo.Session
	repository *Repository
}

type DiscordMessage struct {
	Title       string
	Description string
	Image       string
	Color       string
}

func CreateDiscord(token string, storage *Repository) DiscordClient {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.WithError(err).Fatal("Could not setup discord session")
	}

	return DiscordClient{
		session:    session,
		repository: storage,
	}
}

func (d *DiscordClient) Connect() error {
	err := d.session.Open()
	if err != nil {
		log.WithError(err).Fatal("Could not connect to discord")
	}
	return err
}

func (d *DiscordClient) Close() {
	log.Info("Closing discord session")
	d.session.Close()
}

func (d *DiscordClient) Post(channelID string, game Game) error {
	if len(game.Promotions.PromotionalOffers) < 1 {
		return fmt.Errorf("No promotional offers for %s", game.Title)
	}
	// drop error for now
	if ok, _ := d.IsPublished(channelID, game); ok {
		log.WithFields(log.Fields{"channel": channelID, "game": game.Title}).Info("Game is already published on channel")
		return nil
	}

	imgUrl := game.Images[1].URL

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00ff00, // Green
		Description: game.Description,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Promotion starts:",
				Value:  game.Promotions.PromotionalOffers[0].Offers[0].Start.String(),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Promotion ends:",
				Value:  game.Promotions.PromotionalOffers[0].Offers[0].End.String(),
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: imgUrl,
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     game.Title,
		//To complete - need to concat https://www.epicgames.com/store/en-US/product/ + game.ProductSlug
		//However, it might be better to put either put this in the db directly, or make it work with Variables set in api.go here
		URL: "https://www.epicgames.com/store/en-US/product/" + game.ProductSlug,
	}
	_, err := d.session.ChannelMessageSendEmbed(channelID, embed)

	if err != nil {
		log.WithError(err).Fatal("Unable to post message on discord")
	}
	log.WithFields(log.Fields{"channel": channelID, "game": game.Title}).Info("published game")

	if err := d.MarkAsPublished(channelID, game); err != nil {
		log.WithError(err).Error("Could not mark game as published")
	}

	return err
}

func (d *DiscordClient) MarkAsPublished(channelID string, game Game) error {
	return d.repository.MarkPublished(channelID, game.Title, game.Id, "epic")
}

func (d *DiscordClient) IsPublished(channelID string, game Game) (bool, error) {
	return d.repository.IsPublished(channelID, game.Title, "epic")
}
