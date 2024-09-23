package rabbit

import (
	"errors"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp.Connection
	queueName string
	channel *amqp.Channel
}

type MessageHandler func(amqp.Delivery, string) error

func NewConsumer() (Consumer, error) {
	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s",os.Getenv("rb_user"),os.Getenv("rb_pwd"),os.Getenv("rb_ip"))

	connection, err := Connect(rabbitMQURL)
	if err != nil {
		return Consumer{} , err
	}

	consumer := Consumer {
		conn: connection,
	}
	return consumer, nil

}


func (c *Consumer) SetQueue(queueName string) error {
	channel, err := c.conn.Channel()
	if err != nil {
		return err
	}
	q, err := DeclareQueue(channel, queueName)
	if err != nil {
		return err
	}

	c.queueName = q.Name
	c.channel = channel


	return nil
}

func (c *Consumer) Listen(handler MessageHandler) error {

	if c.queueName == ""{
		return errors.New("no se ha declarado una cola")
	}
	defer func ()  {
		log.Println("Fin del metodo Listen")
		c.channel.Close()
	}()


	messages, err := c.channel.Consume(
		c.queueName, // cola
		"",     // consumidor
		true,   // auto-ack
		false,  // exclusiva
		false,  // no-local
		false,  // sin esperar
		nil,    // argumentos
	)
	if err != nil {
		return err
	}

	go func ()  {
		for d := range messages {
			if err := handler(d, c.queueName); err != nil {
				log.Printf("Error handling message: %s", err)
			}
		}
		defer func ()  {
			log.Println("Se ha terminado la escucha en la cola:",c.queueName)
		}()
	}()
	
	forever := make(chan bool)
	fmt.Printf("Esperando mensajes en [%s]\n", c.queueName)
	<-forever
	
	return nil
}

func (c *Consumer) Consume() (<-chan amqp.Delivery, error){

	messages, err := c.channel.Consume(
		c.queueName, // cola
		"",     // consumidor
		true,   // auto-ack
		false,  // exclusiva
		false,  // no-local
		false,  // sin esperar
		nil,    // argumentos
	)
	if err != nil {
		return nil, err
	}

	return messages, nil

}