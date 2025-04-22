package dto

type GenericoDTO struct {
	Estado  string
	Mensaje string
}

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
