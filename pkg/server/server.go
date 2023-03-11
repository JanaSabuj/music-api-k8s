package server

import (
	"github.com/JanaSabuj/music-api-k8s/pkg/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var gcfg *config.Config // global config for server

func Start(cfg *config.Config) {
	gcfg = cfg
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/music", GetSongHandler).Methods("GET")
	r.HandleFunc("/api/v1/music", PostSongHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
