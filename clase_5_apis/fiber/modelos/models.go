package modelos

type EquipoModel struct {
	Id     int    `json:"id" db:"id"`
	Nombre string `json:"nombre" db:"nombre"`
	Liga   string `json:"liga" db:"liga"`
}

type JugadorModel struct {
	Id       int         `json:"id" db:"id"`
	Nombre   string      `json:"nombre" db:"nombre"`
	Posicion string      `json:"posicion" db:"posicion"`
	EquipoId int         `json:"equipo_id" db:"equipo_id"`
	Equipo   EquipoModel `json:"equipo" db:"equipo"`
}

func (e EquipoModel) Table() string {
	return "equipo"
}

func (j JugadorModel) Table() string {
	return "jugadores"
}
