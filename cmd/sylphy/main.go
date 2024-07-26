package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/olehvolynets/sylphy"
	"github.com/olehvolynets/sylphy/config"
)

var configFilePath = flag.String("config", "sylphy.config.yml", "path to sylphy config file")

func main() {
	flag.Parse()

	configFile, err := os.Open(*configFilePath)
	if err != nil {
		sylphy.FatalMsg("failed to open config file", err)
	}

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		sylphy.FatalMsg("failed to load config", err)
	}
	slog.Debug("Config", "cfg", *cfg)

	app, err := sylphy.NewSylphy(os.Stdout, cfg)
	if err != nil {
		sylphy.Fatal(err)
	}

	// TODO: change to stdio
	f, err := os.Open("tmp/pyra_sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Start(f); err != nil {
		log.Fatal(err)
	}
}
