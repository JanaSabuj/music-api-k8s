package musicapp

type Song struct {
	Id     string `gorm:"PrimaryKey"`
	Name   string
	Artist string
}
