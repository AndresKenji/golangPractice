package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Acá nos conectamos a RabbitMQ o mostramos un mensaje de error en caso que ocurra.
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Falla al conectar con RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Falla al abrir el canal")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"golang-queue", // nombre
		false,          // durable
		false,          // borrar cuando no se use
		false,          // exclusiva
		false,          // sin esperar
		nil,            // argumentos
	)
	failOnError(err, "Falla al declara la cola")

	msgs, err := ch.Consume(
		q.Name, // cola
		"",     // consumidor
		true,   // auto-ack 
		false,  // exclusiva
		false,  // no-local
		false,  // sin esperar
		nil,    // argumentos
	)
	failOnError(err, "Falla al registrar el consumidor")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Se ha recibido el mensaje: %s", d.Body)
		}
	}()

	log.Printf(" [*] Esperando por mensajes. Para salir presione CTRL+C")
	<-forever
}