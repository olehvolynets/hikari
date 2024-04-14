package main

import (
	"flag"
)

var configFilePath = flag.String("config", "sylphy.config.pkl", "path to sylphy config file")

func main() {
	flag.Parse()

	// f, err := os.Open("tmp/tmp_size.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// app := &sylphy.Sylphy{}
	// if err = app.Start(f); err != nil {
	// 	log.Fatal(err)
	// }
}
