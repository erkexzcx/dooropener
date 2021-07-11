package main

import (
	"flag"
	"log"

	"dooropener"
	"dooropener/config"
)

var flagConfigPath = flag.String("config", "dooropener.yml", "path to config file")

func main() {
	flag.Parse()

	c, err := config.NewConfig(*flagConfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	dooropener.Start(c)
}
