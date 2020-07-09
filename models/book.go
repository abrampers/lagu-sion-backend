package models

type Book struct {
	ID        uint `gorm:"primary_key"`
	ShortName string
	LongName  string
}
