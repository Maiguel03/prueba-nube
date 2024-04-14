package db

import (
	"database/sql"
	"fmt"
	"log"
	"prueba/models"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dbHost := "dpg-cnn7jl6n7f5s73da6f70-a.oregon-postgres.render.com"
	dbPort := "5432"
	dbUser := "dek"
	dbPassword := "rubqwjCMeRfKkwvGsGeYudqOpmKtbo34"
	dbName := "empanada"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
		dbHost, dbPort, dbUser, dbName, dbPassword)

	DB, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return DB, nil
}

func CrearUsuario(Usuario models.Usuario) error {
	db, _ := ConnectDB()
	defer db.Close()
	stmt, err := db.Prepare("Insert into Usuario (CI ,Nombre ,Apellido ,Usuario ,Contraseña ,Categoria) values ($1, $2, $3, $4, $5, $6)")

	if err != nil {
		return err
	}
	
	_, err = stmt.Exec(Usuario.CI, Usuario.Nombre, Usuario.Apellido, Usuario.User, Usuario.Pass, "Administrador")

	if err != nil {
		return err
	}

	return nil
}

func Login(Usuario string) (error, string) {
	db, err := ConnectDB()
	

	if err != nil {
		log.Print(err)
		return err, ""
	}
	var HashContraseña string
	err = db.QueryRow("select Contraseña from usuario where Usuario = $1", Usuario).Scan(&HashContraseña)

	if err != nil {
		if err == sql.ErrNoRows{
		log.Print(err)
		return err, ""
		} else{
			log.Print("Error al consulta la base de datos", err)
			return err, ""
		}
	}
	return nil, HashContraseña
}

func ValidarLogin(Usuario interface{}) (error, models.Usuario) {
	db, err := ConnectDB()
	

	if err != nil {
		log.Print(err)
		return err, models.Usuario{}
	}
	var user struct{
		CI string 
		Nombre string 
		Apellido string 
		User string 
		Categoria string
	}
	err = db.QueryRow("select CI ,Nombre ,Apellido ,Usuario ,Categoria from usuario where Usuario = $1", Usuario).Scan(&user.CI, &user.Nombre, &user.Apellido, &user.User, &user.Categoria)

	if err != nil {
		if err == sql.ErrNoRows{
		log.Print(err)
		return err, models.Usuario{}
		} else{
			log.Print("Error al consulta la base de datos", err)
			return err, models.Usuario{}
		}
	}
	return nil, models.Usuario{User: user.User, Nombre: user.Nombre, Apellido: user.Apellido, CI: user.CI}
}
