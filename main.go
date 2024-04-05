package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type usuario struct {
	User string 
	Pass string 
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS") // Especifica los métodos permitidos
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept") // Especifica las cabeceras permitidas
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	if r.Method == "POST" {
		User := usuario{}
		json.NewDecoder(r.Body).Decode(&User)
		json.NewEncoder(w).Encode(User)
		fmt.Println(User)
		return
	}

	if r.Method == "GET" {
		User := usuario{User: "miguel", Pass: "1234"}
		json.NewEncoder(w).Encode(User)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Login)

	server := http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	server.ListenAndServe()
}

