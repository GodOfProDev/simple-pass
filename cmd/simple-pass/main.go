package main

import (
	"github.com/godofprodev/simple-pass/internal/config"
	"github.com/godofprodev/simple-pass/internal/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("there was an issue loading .env")
	}

	serverCfg, err := config.NewServerConfig()
	if err != nil {
		log.Fatal("there was an issue reading the config")
	}

	r := router.New()

	r.RegisterMiddlewares()
	r.RegisterHandlers()

	err = r.Listen(serverCfg)
	if err != nil {
		log.Fatal("there was an issue listening to port 8080: ", err)
	}
}
