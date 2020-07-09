package models

type Song struct {
	ID     uint    `gorm:"primary_key"`
	Number int     `json:"number"`
	Title  string  `json:"title"`
	Verses []Verse `json:"verses"`
	Reff   Verse   `json:"reff"`
	BookID uint    `json:"bookID"`
	Book   Book
}
