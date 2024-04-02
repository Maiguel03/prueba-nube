package main

import (
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := r.FormValue("nombre")
		pass := r.FormValue("contrasenia")

		log.Print("Usuario: ", user, "Contraseña: ", pass)
		fmt.Fprint(w, "Si funcionó jajaja, eres gey si ves esto")
		return
	}
}

func main() {
	http.HandleFunc("/", Login)

	log.Print("Servidor corriendo")
	http.ListenAndServe("localhost:8080", nil)
}
