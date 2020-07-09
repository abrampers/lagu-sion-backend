package models

type Verse struct {
	ID       uint   `gorm:"primary_key"`
	Contents string `json:"contents"`
	SongID   uint
}
