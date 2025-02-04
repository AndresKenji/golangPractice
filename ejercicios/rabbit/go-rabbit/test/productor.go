package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// Aca manejamos la forma en la que los errores son mostrados por consola.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Tomamos el mensaje desde la terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("¿Qué mensaje quieres enviar?")
	mPayload, _ := reader.ReadString('\n')
	// Acá nos conectamos a RabbitMQ o mostramos un mensaje de error en caso que ocurra.
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Falla al conectar con RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Falla al abrir el canal")
	defer ch.Close()

	// Creamos una cola a la cual enviaremos el mensaje.
	q, err := ch.QueueDeclare(
		"golang-queue", // Nombre
		false,          // durable
		false,          // Borrar cuando no se use
		false,          // exclusiva
		false,          // No esperar
		nil,            // argumentos
	)
	failOnError(err, "Falla al declarar la cola")

	// Enviamos la carga o payload como mensaje.
	//body := "Golang es el futuro!"
	err = ch.Publish(
		"",     // intercambio
		q.Name, // llave de enrutamiento
		false,  // obligatorio
		false,  // inmediato
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(mPayload),
		})
	// Si hay algun error publicando el mensaje se mostrara en la consola.
	failOnError(err, "Falla publicando el mensaje")
	log.Printf(" [x] Se ha enviado el mensaje: %s \n", mPayload)
}
