package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func messageHandle(data amqp.Delivery, qname string) error {
	var message Message
	err := json.Unmarshal(data.Body, &message)
	if err != nil {
		return err
	}
	switch qname {

	case "urgente":

		log.Printf("Se ha recibido un mensaje urgente del usuario %s \nMensaje: %s", message.Name, message.Msg)

	case "info":

		log.Printf("Se ha recibido información del usuario %s \nInfo: %s", message.Name, message.Msg)

	case "error":

		log.Printf("Se ha presentado un error con el usuario %s \nError: %s", message.Name, message.Msg)

	}

	return nil
}
