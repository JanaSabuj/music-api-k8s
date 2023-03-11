package server

import (
	"encoding/json"
	"github.com/JanaSabuj/music-api-k8s/pkg/musicapp"
	"log"
	"net/http"
)

func GetSongHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GET /api/v1/music")

	// get songs
	songs, err := musicapp.GetAllSongs(gcfg)
	if err != nil {
		log.Fatal("get songs api failed", err.Error())
	}

	// response
	json.NewEncoder(w).Encode(songs)
}
