package dto

type CategoriaDTO struct {
	Nombre string `json:"nombre"`
}

type CanalDTO struct {
	Nombre    string `json:"nombre"`
	UsuarioId uint   `json:"usuario_id"`
}
