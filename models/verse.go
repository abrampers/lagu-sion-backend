package models

import (
	"github.com/abrampers/lagu-sion-backend/lagusion"
	"strings"
)

type Verse struct {
	Id       uint   `gorm:"primary_key"`
	Contents string `json:"contents"`
	SongID   uint
}

func (v Verse) LaguSionVerse() *lagusion.Verse {
	return &lagusion.Verse{Contents: strings.FieldsFunc(v.Contents, func(c rune) bool {
		return c == '\n'
	})}
}
