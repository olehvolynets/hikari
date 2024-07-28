package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/olehvolynets/hikari"
	"github.com/olehvolynets/hikari/config"
)

var configFilePath = flag.String("config", "hikari.config.yml", "path to hikari config file")

func main() {
	flag.Parse()

	configFile, err := os.Open(*configFilePath)
	if err != nil {
		hikari.FatalMsg("failed to open config file", err)
	}

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		hikari.FatalMsg("failed to load config", err)
	}
	// slog.Debug("Config", "cfg", *cfg)

	app, err := hikari.NewHikari(os.Stdout, cfg)
	if err != nil {
		hikari.Fatal(err)
	}

	if err = app.Start(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
