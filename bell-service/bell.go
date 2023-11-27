package main

import (
	"bell-service/config"
	"bell-service/internal/app"
	"log"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
	}
	app.Run(cfg)
}
