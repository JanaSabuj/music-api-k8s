package musicapp

import (
	"gorm.io/gorm"
	"log"
)

type musicDaoInterface interface {
	Get(db *gorm.DB) ([]*Song, error)
}

type musicApiOrm struct {
}

func NewMusicOrm() musicDaoInterface {
	return &musicApiOrm{}
}

func (m *musicApiOrm) Get(db *gorm.DB) ([]*Song, error) {
	var songs []*Song
	if err := db.Find(&songs).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("%d rows found", len(songs))

	return songs, nil
}
