package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	api := Api{}
	config := LoadConfig()
	repo, err := CreateRepository(config.Storage)
	if err != nil {
		log.WithError(err).Fatal("could not open storage")
		os.Exit(1)
	}
	client := CreateDiscord(config.ApiKey, repo)


	ctx, done := context.WithCancel(context.Background())
	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		client.Connect()
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
		select {
		case sig := <-signalChannel:
			log.Debug("Received signal: %s\n", sig)
			done()
		case <-gctx.Done():
			log.Debug("closing goroutine")
			client.Close()
			return gctx.Err()
		}

		return nil
	})

	// Epic gamestore
	g.Go(func() error {
		ticker := time.NewTicker(time.Duration(config.Interval) * time.Second)

		for {
			select {
			case <-ticker.C:
				response, err := api.get()
				if err != nil {
					log.WithError(err).Fatal("error while fetching games from api")
				}
				for _, game := range response.Data.Catalog.CatalogOffers.Elements {
					err := client.Post(config.ChannelID, game)
					if err != nil {
						return err
					}
				}

				log.Debug("finished fetching free games")
			case <-gctx.Done():
				log.Debug("stopping loop")
				return gctx.Err()
			}
		}
	})

	if err := g.Wait(); err == nil || err == context.Canceled {
		log.Info("shutting down...")
	} else {
		log.WithError(err).Fatal("received error")
	}
	client.Close()
	os.Exit(1)
}
