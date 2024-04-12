package main

import (
	"prueba/handlers"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.POST("/registro", handlers.Registro) 
	r.POST("/login", handlers.Login)
	r.GET("/validar", handlers.Validar)
	r.Run(":8080")

}
