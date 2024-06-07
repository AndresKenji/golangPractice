package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Iniciando el servicio de autenticación")

	// Conexión a la db
	conn := connectToDB()
	if conn == nil {
		log.Panic("No se ha logrado conexión a Postgres!")
	}

	// set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {

			log.Println("Postgres no esta listo aun...")
			counts++
		} else {
			log.Println("Conectado a Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}
		log.Println("Esperando por dos segundos...")
		time.Sleep(2 * time.Second)
		continue
	}
}
