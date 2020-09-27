package models

import (
	"github.com/abrampers/lagu-sion-backend/lagusion"
	"github.com/google/uuid"
)

var LaguSionBook Book = Book{Base: Base{Id: uuid.New()}, ShortName: "LS", LongName: "Lagu Sion"}
var LaguSionEdisiLengkapBook Book = Book{Base: Base{Id: uuid.New()}, ShortName: "LSEL", LongName: "Lagu Sion Edisi Lengkap"}

type Book struct {
	Base
	ShortName string
	LongName  string
}

func GetBook(id uint32) Book {
	if id == 1 {
		return LaguSionBook
	} else if id == 2 {
		return LaguSionEdisiLengkapBook
	} else {
		return LaguSionBook
	}
}

func LaguSionSongBook(id uint32) *lagusion.Book {
	if id == 1 {
		return LaguSionBook.LaguSionSongBook()
	} else if id == 2 {
		return LaguSionEdisiLengkapBook.LaguSionSongBook()
	} else {
		return nil
	}
}

func (b Book) LaguSionSongBook() *lagusion.Book {
	return &lagusion.Book{
		Id:        &lagusion.UUID{Value: b.Id.String()},
		ShortName: b.ShortName,
		LongName:  b.LongName,
	}
}
