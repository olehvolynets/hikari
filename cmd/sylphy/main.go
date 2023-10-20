package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olehvolynets/sylphy"

	_ "github.com/fatih/color"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	confFile, err := os.ReadFile("sample_config.json")
	check(err)

	cfg, err := sylphy.NewConfig(confFile)
	check(err)

	marshalled, _ := json.MarshalIndent(cfg, "", "   ")
	fmt.Println(string(marshalled))
}
