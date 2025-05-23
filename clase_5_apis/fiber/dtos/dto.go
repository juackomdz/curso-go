package dtos

type EquiposDTO struct {
	Nombre string `json:"nombre" db:"nombre"`
	Liga   string `json:"liga" db:"liga"`
}

type JugadoresDTO struct {
	Nombre   string `json:"nombre" db:"nombre"`
	Posicion string `json:"posicion" db:"posicion"`
	EquipoId int    `json:"equipo_id" db:"equipo_id"`
}
