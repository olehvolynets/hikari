package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/olehvolynets/sylphy"
	"github.com/olehvolynets/sylphy/scheme"
)

var configFilePath = flag.String("config", ".sylphy.config.yml", "path to sylphy config file")

func check(err error) {
	if err != nil {
		panic(fmt.Errorf("sylphy: %w", err))
	}
}

func main() {
	flag.Parse()

	confFile, err := os.ReadFile(*configFilePath)
	check(err)

	scheme, err := scheme.New(confFile)
	check(err)

	s := sylphy.New(scheme)

	payloadFile, err := os.ReadFile("sample.json")
	check(err)

	payload := make([]map[string]any, 0)
	err = json.Unmarshal(payloadFile, &payload)
	check(err)

	for _, row := range payload {
		s.Print(row)
	}
}
