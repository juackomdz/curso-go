package modelos

import "github.com/uptrace/bun"

type TematicaModel struct {
	bun.BaseModel `bun:"table:tematica,alias:t"`

	Id     int    `bun:",pk,autoincrement" json:"id"`
	Nombre string `bun:"nombre,notnull" json:"nombre"`
}

type PeliculaModel struct {
	bun.BaseModel `bun:"table:peliculas,alias:p"`

	Id          int           `bun:",pk,autoincrement" json:"id"`
	Nombre      string        `bun:"nombre,notnull" json:"nombre"`
	Descripcion string        `bun:"descripcion" json:"descripcion"`
	TematicaId  int           `bun:"tematica_id" json:"tematica_id"`
	Tematica    TematicaModel `bun:"rel:belongs-to,join:tematica_id=id" json:"tematica"`
}
