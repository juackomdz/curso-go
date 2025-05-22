package modelos

type EquipoModel struct {
	Id     int    `json:"id" db:"id"`
	Nombre string `json:"nombre" db:"nombre"`
	Liga   string `json:"liga" db:"liga"`
}

func (e EquipoModel) Table() string {
	return "equipo"
}
