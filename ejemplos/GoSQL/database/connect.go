package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Abrir conexión a una db
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	// Verificar conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Conexión a Mysql desde Go ok!")

	return db, nil

}
