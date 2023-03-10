package main

import (
	"fmt"
	"github.com/JanaSabuj/music-api-k8s/pkg/config"
	"github.com/JanaSabuj/music-api-k8s/pkg/musicapp"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to music api...")

	cfg, err := config.NewConfig()
	fmt.Println(cfg, err)
	if err != nil {
		panic(err)
	}

	// migrate evergreen_music_db
	if err := musicapp.DbInit(cfg.DB); err != nil {
		panic(err)
	}

	// router
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/music", getBookHandler).Methods("GET")
}
