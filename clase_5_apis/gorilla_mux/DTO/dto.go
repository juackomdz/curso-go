package dto

type CategoriaDTO struct {
	Nombre string `json:"nombre"`
}

type CanalDTO struct {
	Nombre    string `json:"nombre"`
	UsuarioId uint   `json:"usuario_id"`
}

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	PerfilId uint   `json:"perfil_id"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	User  string `json:"user"`
	Token string `json:"token"`
}
