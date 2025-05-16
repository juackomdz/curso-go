package dto

import "github.com/uptrace/bun"

type TematicaDTO struct {
	bun.BaseModel `bun:"table:tematica"`
	Nombre        string `json:"nombre"`
}
