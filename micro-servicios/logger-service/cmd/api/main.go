package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// conectar con mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	// crear un context para manejar la desconección

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// cerrar conexión
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// Registrar el servidor RPC
	err = rpc.Register(new(RPCServer))
	go app.rpcListen()

	go app.gRPCListen()

	// iniciar el servidor web
	log.Println("Iniciando servidor en el puerto:", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func (app *Config) rpcListen() error {
	log.Println("Iniciando servidor RPC en puesto:", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}

func connectToMongo() (*mongo.Client, error) {
	// crear opciones de conexión
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// conexión
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error en la conexión:", err)
		return nil, err
	}

	log.Println("Conectado a mongo!")

	return c, nil
}
