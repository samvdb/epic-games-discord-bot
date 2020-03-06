package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	ApiKey    string
	Interval  int64 // in seconds
	ChannelID string
	Storage   string
}

func LoadConfig() Config {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)

	var configuration Config

	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file config.yml")
	}

	// Set undefined variables
	viper.SetDefault("STORAGE", "")
	viper.SetDefault("APIKEY", "")
	viper.SetDefault("CHANNELID", "")
	viper.SetDefault("INTERVAL", "120")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Discord api key is\t", configuration.ApiKey)
	fmt.Println("Discord ChannelID is\t", configuration.ChannelID)
	fmt.Println("Interval is\t", configuration.Interval)
	fmt.Println("Storage path is\t", configuration.Storage)
	return configuration
}
