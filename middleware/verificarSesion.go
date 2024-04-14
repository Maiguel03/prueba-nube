package middleware

import (
	"fmt"
	"log"
	"net/http"
	"prueba/db"
	"prueba/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuntenticarSesion(c *gin.Context) {
	//Obtener cookie
	tokenString, err := c.Cookie("Autorizacion")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Decodificar y validar

	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("empanada"), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//Chequear fecha de vencimiento
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Buscar usuario con el token
		var user models.Usuario
		UserName := claims["sub"]
		err, user = db.ValidarLogin(UserName)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Almacenar datos de usuario para que sean accedibles desde la funcion que la necesite
		c.Set("user", user)

		//Continuar
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
