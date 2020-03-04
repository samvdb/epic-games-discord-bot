package main

import (
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"time"
)

type DiscordClient struct {
	session     *discordgo.Session
	repository *Repository
}

func CreateDiscord(token string, storage *Repository) DiscordClient {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.WithError(err).Fatal("Could not setup discord session")
	}

	return DiscordClient{
		session:     session,
		repository: storage,
	}
}


func (d *DiscordClient) Post(channelID string, game Game) error {

	if d.IsPublished(channelID, game) {
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
				Name:   "Title",
				Value:  game.Title,
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: imgUrl,
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     game.Title,
	}

	//_, err := d.session.ChannelMessage(channelID, "abc")
	//
	//if err != nil {
	//	log.WithError(err).Fatal("Unable to post message on discord")
	//}
	err := d.session.Open()
	if err != nil {
		log.WithError(err).Fatal("Could not connect to discord")
		return err
	}
	defer d.session.Close()
	_, err = d.session.ChannelMessageSendEmbed(channelID, embed)

	if err != nil {
		log.WithError(err).Fatal("Unable to post message on discord")
	}
	log.WithFields(log.Fields{"channel": channelID, "game": game.Title}).Info("published game")

	d.MarkAsPublished(channelID, game)
	return err
}

func (d *DiscordClient) MarkAsPublished(channelID string, game Game) {
	d.repository.MarkPublished(channelID, game)
}

func (d *DiscordClient) IsPublished(channelID string, game Game) bool {
	return d.repository.IsPublished(channelID, game)
}
