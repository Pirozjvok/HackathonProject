package main

import (
	"audio-saver-service/config"
	"audio-saver-service/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
	}

	app.Run(cfg)
}
