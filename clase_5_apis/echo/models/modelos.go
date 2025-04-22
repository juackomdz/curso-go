package models

type Categoria struct {
	Id     string `json:"id"`
	Nombre string `json:"nombre"`
}

type Producto struct {
	Id          string `json:"id"`
	Nombre      string `json:"nombre"`
	Precio      int    `json:"precio"`
	Stock       int    `json:"stock"`
	CategoriaId string `json:"categoria_id"`
}
