package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)

	server := NewAPIServer(":3000", store)
	server.Serve()
}

// https://www.youtube.com/watch?v=2JNUmzuBNV0&list=PL5I1bZZq1Y9yjLHR72TS9xMkt6_2vXAWy&index=18