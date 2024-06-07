package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

// RPCServer es el tipo para nuestro servidor RPC. Los Metodos que lo toman como receiver estan disponibles
// sobre RPC, mientras sean exportados.
type RPCServer struct{}

// RPCPayload es el tipo para recibir nuestros datos desde RPC
type RPCPayload struct {
	Name string
	Data string
}

// LogInfo Escribe la información del payload en mongo
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	// resp es el mensaja que se envia al RPC caller
	*resp = "Processed payload via RPC:" + payload.Name
	return nil
}
