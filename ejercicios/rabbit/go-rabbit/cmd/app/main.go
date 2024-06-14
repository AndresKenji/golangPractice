package main

import "rabbitgo/internal/rabbit"

var colas = []string{"info",
					"error",
					"urgente",
					}

func main() {

	for _, queue := range colas {

		consumer, err := rabbit.NewConsumer()
		failOnError(err, "Error al conectar con RabbitMQ")
		err = consumer.SetQueue(queue)
		failOnError(err, "Falla al declarar la cola")
		go func ()  {
			err = consumer.Listen(messageHandle)
			failOnError(err, "Falla al escuchar la cola")
		}()
	}


	forever := make(chan bool)
	<-forever

}