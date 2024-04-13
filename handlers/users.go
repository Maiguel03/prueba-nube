package handlers

import (
	"log"
	"net/http"
	"prueba/db"
	"prueba/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Registro(c *gin.Context) {
	//Obtener datos del cliente
	var Body models.Usuario

	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	//Crear hash de la contraseña
	passhash, err := bcrypt.GenerateFromPassword([]byte(Body.Pass), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error al crear hash de la contraseña",
		})
		return
	}

	//Crear usuario
	User := models.Usuario{
		User: Body.User,
		Pass: string(passhash), 
		Nombre: Body.Nombre,
		Apellido: Body.Apellido,
		CI: Body.CI,
	}

	if err = db.CrearUsuario(User); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error al crear el usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	//Obtener usuario y contraseña del cliente
	var User struct{
		User string `json:"userName" bindin:"required"`
		Pass string	`json:"password" bindin:"required"`
	}

	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}

	//Consultar con base de datos si existe un usuario llamado asi y retonar el hash
	err, passhash := db.Login(User.User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Usuario o contraseña invalidos",
		})
		return
	}
	
	//Comparar contraseña enviada por el usuario y contraseña almacenada en la bd
	err = bcrypt.CompareHashAndPassword([]byte(passhash), []byte(User.Pass))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Usuario o contraseña invalidos",
		})
		return
	}

	//Crear token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": User.User,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// Iniciar sesion y obtener el token completo como un string usando
	tokenString, err := token.SignedString([]byte("empanada"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error al crear token",
		})
		return
	}

	//Enviar respuesta
	c.SetCookie("Autorizacion", tokenString, 3600*24*7, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"Token": tokenString,
	})
}

func Validar (c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Mensaje": "Sesión válida",
	})

}
