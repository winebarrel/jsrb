package main

import (
	"log"
	"os"

	"github.com/winebarrel/jsrt"
)

func init() {
	log.SetFlags(0)
}

func main() {
	file, key := parseArgs()
	defer file.Close()

	err := jsrt.Sort(file, key, os.Stdout)

	if err != nil {
		log.Fatal(err)
	}
}
