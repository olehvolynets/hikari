package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// var configFilePath = flag.String("config", ".sylphy.config.yml", "path to sylphy config file")

func main() {
	f, err := os.Open("tmp_size.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)

	dec := json.NewDecoder(r)

	for {
		v := make(map[string]any)

		if err := dec.Decode(&v); err == io.EOF {
			break
		} else if err != nil {
			mr := io.MultiReader(dec.Buffered(), r)
			r = bufio.NewReader(mr)

			line, _ := r.ReadString('\n')
			fmt.Print(line)
			dec = json.NewDecoder(r)
			continue
		}

		fmt.Println(v)
	}
}
