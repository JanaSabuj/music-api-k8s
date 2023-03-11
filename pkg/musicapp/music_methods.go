package musicapp

import (
	"github.com/JanaSabuj/music-api-k8s/pkg/config"
	"log"
)

var musicOrm = NewMusicOrm()

func GetAllSongs(cfg *config.Config) ([]*Song, error) {
	// get songs
	songs, err := musicOrm.Get(cfg.DB)
	if err != nil {
		log.Fatal("Error getting songs ", err.Error())
	}

	return songs, err
}
