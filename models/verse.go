package models

import (
	"github.com/abrampers/lagu-sion-backend/lagusion"
	"github.com/google/uuid"
)

type Verse struct {
	Base
	Contents string `json:"contents"`
	VersesOf uuid.UUID
	ReffOf   uuid.UUID
}

func (v Verse) LaguSionVerse() *lagusion.Verse {
	return &lagusion.Verse{Contents: v.Contents}
}
