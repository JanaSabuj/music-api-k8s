package musicapp

import (
	"gorm.io/gorm"
	"log"
)

const (
	CreateBatchSizeDefault = 100
)

type musicDaoInterface interface {
	Get(db *gorm.DB) ([]*Song, error)
	Post(db *gorm.DB, song *Song) error
}

type musicApiOrm struct {
}

func NewMusicOrm() musicDaoInterface {
	return &musicApiOrm{}
}

func (m *musicApiOrm) Get(db *gorm.DB) ([]*Song, error) {
	songsArr := make([]*Song, 0)
	if err := db.Model(&Song{}).Find(&songsArr).Error; err != nil {
		return nil, err
	}
	log.Printf("%d rows found", len(songsArr))

	return songsArr, nil
}

func (m *musicApiOrm) Post(db *gorm.DB, song *Song) error {
	if err := db.Session(&gorm.Session{FullSaveAssociations: true, CreateBatchSize: CreateBatchSizeDefault}).Model(Song{}).Create(song).Error; err != nil {
		return err
	}

	return nil
}
