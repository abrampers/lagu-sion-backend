package models

import (
	"github.com/abrampers/lagu-sion-backend/lagusion"
	"log"
)

type Book struct {
	Id        uint32 `gorm:"primary_key"`
	ShortName string
	LongName  string
}

func LaguSionSongBook(id uint32) lagusion.SongBook {
	if id == uint32(lagusion.SongBook_ALL) {
		return lagusion.SongBook_ALL
	} else if id == uint32(lagusion.SongBook_LAGU_SION) {
		return lagusion.SongBook_LAGU_SION
	} else if id == uint32(lagusion.SongBook_LAGU_SION_EDISI_LENGKAP) {
		return lagusion.SongBook_LAGU_SION_EDISI_LENGKAP
	} else {
		log.Println("Invalid SongBook enum")
		return lagusion.SongBook_ALL
	}
}

func (b Book) LaguSionSongBook() lagusion.SongBook {
	if b.Id == uint32(lagusion.SongBook_ALL) {
		return lagusion.SongBook_ALL
	} else if b.Id == uint32(lagusion.SongBook_LAGU_SION) {
		return lagusion.SongBook_LAGU_SION
	} else if b.Id == uint32(lagusion.SongBook_LAGU_SION_EDISI_LENGKAP) {
		return lagusion.SongBook_LAGU_SION_EDISI_LENGKAP
	} else {
		log.Println("Invalid SongBook enum")
		return lagusion.SongBook_ALL
	}
}
