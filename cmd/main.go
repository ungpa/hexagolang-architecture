package main

import (
	"log"

	"github.com/ungpa/goon-ah-lang/internal/app"
	"github.com/ungpa/goon-ah-lang/internal/infrastructure/config"
)

func main() {
	cfg, err := config.NewAppConfig("./config/config.json")
	if err != nil {
		log.Fatalln("Unable to read configuration file", err)
	}

	application := app.NewApp(cfg)
	application.Run()

}
