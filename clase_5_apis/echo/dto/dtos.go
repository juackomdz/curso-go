package dto

type ProductoDTO struct {
	Nombre      string `json:"nombre"`
	Precio      int    `json:"precio"`
	Descripcion string `json:"descripcion"`
	Stock       int    `json:"stock"`
	CategoriaId string `json:"categoria_id"`
}

type CategoriaDTO struct {
	Nombre string `json:"nombre"`
}

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Correo   string `json:"correo"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
