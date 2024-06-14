package rabbit

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	conn *amqp.Connection
}

func NewProducer() (Producer, error) {
	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s",os.Getenv("rb_user"),os.Getenv("rb_pwd"),os.Getenv("rb_ip"))

	connection, err := Connect(rabbitMQURL)
	if err != nil {
		return Producer{} , err
	}

	producer := Producer {
		conn: connection,
	}

	return producer, nil

}

func (p *Producer) Push(message []byte, exchange string) error{
	channel, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	err = channel.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: message,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) CloseConn(){
	p.conn.Close()
}

func (p *Producer) BindQueues(exchange string, queues []string) error {
	channel, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	err = DeclareExchange(channel,exchange, "fanout")
	if err != nil {
		return err
	}
	for _, queue := range queues {
		// Creamos una cola a la cual enviaremos el mensaje. (Si ya existe no pasa nada 😎)
		q, err := DeclareQueue(channel,queue)
		if err != nil {
			return err
		}
		log.Printf("Vinculando cola %s al exchange %s",queue, exchange)
		err = channel.QueueBind(
			q.Name,
			"",
			exchange,
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}

	return nil
}