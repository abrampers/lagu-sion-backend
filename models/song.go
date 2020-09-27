package models

import (
	"encoding/json"
	"github.com/abrampers/lagu-sion-backend/lagusion"
	"github.com/google/uuid"
)

type Song struct {
	Base
	Number     uint32  `json:"number"`
	Title      string  `json:"title"`
	Verses     []Verse `json:"verses" gorm:"foreignkey:VersesOf"`
	Reff       Verse   `json:"reff" gorm:"foreignkey:ReffOf"`
	BookNumber uint32  `json:"bookNumber"`
	BookId     uuid.UUID
	Book       Book `gorm:"PRELOAD:false"`
}

func NewSong(jsonBlob []byte) (*Song, error) {
	var song Song
	if err := json.Unmarshal(jsonBlob, &song); err != nil {
		return nil, err
	}
	song.Book = GetBook(song.BookNumber)
	return &song, nil
}

func (s Song) LaguSionSong() *lagusion.Song {
	var verses []*lagusion.Verse
	for _, verse := range s.Verses {
		verses = append(verses, verse.LaguSionVerse())
	}
	return &lagusion.Song{
		Id:     &lagusion.UUID{Value: s.Id.String()},
		Number: s.Number,
		Title:  s.Title,
		Verses: verses,
		Reff:   s.Reff.LaguSionVerse(),
		Book:   LaguSionSongBook(s.BookNumber),
	}
}
