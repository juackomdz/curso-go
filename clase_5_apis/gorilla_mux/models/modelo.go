package models

import db "clase_5/database"

type Usuario struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Nombre   string `gorm:"type:varchar(100)" json:"nombre"`
	Apellido string `gorm:"type:varchar(100)" json:"apellido"`
	Email    string `gorm:"type:varchar(20)" json:"email"`
}

type Usuarios []Usuario

func Migrar() {
	db.Conectar().AutoMigrate(&Usuario{})
}
