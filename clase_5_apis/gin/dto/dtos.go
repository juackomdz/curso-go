package dto

import "github.com/uptrace/bun"

type TematicaDTO struct {
	bun.BaseModel `bun:"table:tematica"`
	Nombre        string `json:"nombre"`
}

type PeliculaDTO struct {
	bun.BaseModel `bun:"table:peliculas"`

	Nombre      string `bun:"nombre,notnull" json:"nombre"`
	Descripcion string `bun:"descripcion" json:"descripcion"`
	TematicaId  int    `bun:"tematica_id" json:"tematica_id"`
}
