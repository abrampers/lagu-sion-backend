package models

import (
	"encoding/json"
	"github.com/abrampers/lagu-sion-backend/lagusion"
)

type Song struct {
	Id     uint32  `gorm:"primary_key"`
	Number int32   `json:"number"`
	Title  string  `json:"title"`
	Verses []Verse `json:"verses"`
	Reff   Verse   `json:"reff"`
	BookId uint32  `json:"bookID"`
	Book   Book    `gorm:"PRELOAD:false"`
}

func NewSong(jsonBlob []byte) (*Song, error) {
	var song Song
	if err := json.Unmarshal(jsonBlob, &song); err != nil {
		return nil, err
	}
	return &song, nil
}

func (s Song) LaguSionSong() *lagusion.Song {
	var verses []*lagusion.Verse
	for _, verse := range s.Verses {
		verses = append(verses, verse.LaguSionVerse())
	}
	return &lagusion.Song{
		Id:       s.Id,
		Number:   s.Number,
		Title:    s.Title,
		Verses:   verses,
		Reff:     s.Reff.LaguSionVerse(),
		SongBook: LaguSionSongBook(s.BookId),
	}
}
