package main

import (
	"prueba/handlers"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware es un middleware personalizado para manejar CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// Si el método es OPTIONS, retornar sin llamar a ningún otro middleware o controlador
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Llamar al siguiente middleware o controlador
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Usar el middleware personalizado de CORS
    	r.Use(CORSMiddleware())
	
	r.POST("/registro", handlers.Registro) 
	r.POST("/login", handlers.Login)
	r.GET("/validar", handlers.Validar)
	r.Run(":8080")

}
