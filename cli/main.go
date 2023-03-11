package main

import (
	"github.com/JanaSabuj/music-api-k8s/pkg/config"
	"github.com/JanaSabuj/music-api-k8s/pkg/musicapp"
	"github.com/JanaSabuj/music-api-k8s/pkg/server"
	"log"
)

var cfg *config.Config
var err error

func init() {
	log.Print("Welcome to music api...")

	// get a config
	cfg, err = config.NewConfig()
	if err != nil {
		panic(err)
	}

	// migrate db
	if err = musicapp.DbInit(cfg.DB); err != nil {
		panic(err)
	}
}

func main() {
	server.Start(cfg)
}
