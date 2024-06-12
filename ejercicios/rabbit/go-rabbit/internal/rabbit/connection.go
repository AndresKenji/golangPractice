package rabbit

import (
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(url string) (*amqp.Connection, error) {
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


func DeclareExchange(ch *amqp.Channel, name, _type string) error {
	return ch.ExchangeDeclare(
		name, // name ?
		_type,      // type ?
		true,         // durable ?
		false,        // auto-delete ?
		false,        // is internal ?
		false,        // no-wait
		nil,          // arguments =
	)
}

func DeclareQueue(ch *amqp.Channel,name string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no wait
		nil,   // arguments
	)
}
