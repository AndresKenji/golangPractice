package main

import (
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// intentar conectarse a rabbit
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// escuchar por mensajes

	log.Println("Escuchando y consumiendo mensajes de RabbitMQ...")

	// crear un consumidor
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// observar la cola y consumir eventos
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// no continuar hasta que rabbit esté listo

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			log.Println("Rabbitmq no está listo aun...")
			counts++
		} else {
			log.Println("Connectado a Rabbitmq!")
			connection = c
			break
		}
		if counts > 5 {
			log.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Esperando ...")
		time.Sleep(backOff)
	}

	return connection, nil
}
