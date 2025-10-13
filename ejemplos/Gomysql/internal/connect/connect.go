package connect

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB
// Connect: función para conectar a la base de datos
func Connect() {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	connection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	Db = connection
}

// CloseConnection: función para cerrar la conexión a la base de datos
func CloseConnection() {
	Db.Close()
}



//go get github.com/joho/godotenv
//go get github.com/go-sql-driver/mysql
/*
CREATE TABLE clientes (
id int NOT NULL AUTO_INCREMENT,
nombre varchar(100) NOT NULL,
correo varchar(100) NOT NULL,
telefono varchar(20) NOT NULL,
fecha datetime NOT NULL,
PRIMARY KEY (id)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/



