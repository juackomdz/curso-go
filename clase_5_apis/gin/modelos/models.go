package modelos

import "github.com/uptrace/bun"

type TematicaModel struct {
	bun.BaseModel `bun:"table:tematica"`

	Id     int    `bun:",pk,autoincrement" json:"id"`
	Nombre string `bun:"nombre,notnull" json:"nombre"`
}
