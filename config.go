package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	ApiKey   string
	Interval int64 // in seconds
	ChannelID string
	Storage string
}

func LoadConfig() Config {
	var C Config
	C.ApiKey = os.Getenv("DISCORD_APIKEY")
	C.ChannelID = os.Getenv("DISCORD_CHANNELID")
	C.Storage = os.Getenv("STORAGE")
	if C.Storage == "" {
		C.Storage = "./"
	}
	C.Interval, _ = strconv.ParseInt(os.Getenv("DISCORD_INTERVAL"), 10, 64)
	if C.Interval == 0 {
		C.Interval = 120
	}
	fmt.Println("Discord api key is\t", C.ApiKey)
	fmt.Println("Discord ChannelID is\t", C.ChannelID)
	fmt.Println("Interval is\t", C.Interval)
	return C
}
