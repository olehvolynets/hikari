package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	"github.com/olehvolynets/sylphy"
	"github.com/olehvolynets/sylphy/scheme"
)

var configFilePath = flag.String("config", ".sylphy.config.yml", "path to sylphy config file")

func check(err error) {
	if err != nil {
		log.Fatalf("sylphy: %w", err)
	}
}

func main() {
	flag.Parse()

	confFile, err := os.ReadFile(*configFilePath)
	check(err)

	scheme, err := scheme.New(confFile)
	check(err)

	app := sylphy.New(scheme)

	if len(flag.Args()) == 0 {
		decoder := json.NewDecoder(os.Stdin)
		err := decodingLoop(app, decoder)
		check(err)
	} else {
		for _, fileName := range flag.Args() {
			f, err := os.Open(fileName)
			check(err)

			decoder := json.NewDecoder(f)

			decodingLoop(app, decoder)
		}
	}
}

func decodingLoop(app *sylphy.Sylphy, decoder *json.Decoder) error {
	for {
		payload := make(map[string]any)

		if err := decoder.Decode(&payload); err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		app.Print(payload)
	}

	return nil
}
