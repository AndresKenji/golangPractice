package rabbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}


func NewConsumer(conn *amqp.Connection, queuename string) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
		queueName: queuename,
	}

	return consumer, nil

}



func (consumer *Consumer) ListenTopics(topics []string, exchangeName string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declarando la cola
	q, err := DeclareQueue(ch, consumer.queueName)
	if err != nil {
		return err
	}
	// se agregan los topics al exchange
	for _, s := range topics {
		ch.QueueBind(
			q.Name, // Name
			s, // key 
			exchangeName, // exchange
			false, // no wait
			nil, // args
		)
		if err != nil {
			return err
		}
	}
	// se recuperan los mensajes 
	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	// Se envian los mensajes a su propia gorrutina y se trabajan desde allá
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)
			go handlePayload(payload)
			}
		}()
			
	fmt.Printf("Esperando mensajes en [Exchange, Queue] [logs_topic, %s]\n", q.Name)
	// canal para mantener la función corriendo
	forever := make(chan bool)
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	switch payload.Name {
	case "log", "event":
		// enviar a logs
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}

	case "auth":
		// autenticarse

	// se pueden tener tantos como se necesiten en la lógica
	default:
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}
	}
}

func logEvent(entry Payload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	//log.Println("Entrando al servicio de log con la data: ",string(jsonData))

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}