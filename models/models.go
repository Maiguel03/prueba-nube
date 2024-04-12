package models

type Usuario struct {
	User string `json:"userName" binding:"required"`
	Pass string `json:"password" binding:"required"`
	Nombre string `json:"firstName" binding:"required"`
	Apellido string `json:"lastName" binding:"required"` 
	CI string `json:"dni" binding:"required"`
}
