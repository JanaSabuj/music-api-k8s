package musicapp

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Name   string `json:"name"`
	Artist string `json:"artist"`
}
