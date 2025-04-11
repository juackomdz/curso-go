package models

import (
	db "clase_5/database"
	"time"
)

type Usuario struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Nombre   string `gorm:"type:varchar(100)" json:"nombre"`
	Apellido string `gorm:"type:varchar(100)" json:"apellido"`
	Email    string `gorm:"type:varchar(20)" json:"email"`
}

type Canal struct {
	Id        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre    string  `gorm:"type:varchar(100)" json:"nombre"`
	UsuarioId uint    `json:"usuario_id"`
	Usuario   Usuario `json:"usuario"`
}

type Perfil struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	Nombre string `json:"nombre"`
}

type User struct {
	Id       uint      `gorm:"primaryKey" json:"id"`
	Username string    `gorm:"type:varchar(100)" json:"username"`
	Password string    `gorm:"type:varchar(100)" json:"password"`
	Email    string    `gorm:"type:varchar(20)" json:"email"`
	Fecha    time.Time `json:"fecha"`
	PerfilId uint      `json:"perfil_id"`
	Perfil   Perfil    `json:"perfil"`
}

type Usuarios []Usuario
type Canales []Canal
type Perfiles []Perfil
type Users []User

func Migrar() {
	//db.Conectar().AutoMigrate(&Usuario{})
	db.Conectar().AutoMigrate(&Perfil{}, User{})
}
