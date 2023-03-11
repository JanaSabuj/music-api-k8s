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
		log.Print("GET songs api failed", err.Error())
		w.Write([]byte("Error getting songs!"))
		return
	}

	// response
	json.NewEncoder(w).Encode(songs)
}

func PostSongHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("POST /api/v1/music")

	// post songs
	newsong := musicapp.Song{}
	json.NewDecoder(r.Body).Decode(&newsong)

	err := musicapp.PostSong(gcfg, &newsong)
	if err != nil {
		log.Print("POST songs api failed", err.Error())
		w.Write([]byte("Error posting song !"))
		return
	}

	// response
	w.Write([]byte("New song published !"))
}
