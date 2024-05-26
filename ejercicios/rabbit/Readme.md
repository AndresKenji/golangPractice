# Rabbit MQ & Go

En este proyecto usaremos RabitMQ con docker compose y crearemos un productor y un consumidor usando go, para correr esto en su maquina asegurese de tener instalado lo siguiente:

- Docker
- Docker compose
- Go

<br>

## Rabbitmq

[RabbitMQ][rabbitmq] es un software de intermediación de mensajes (message broker) de código abierto que facilita la comunicación entre aplicaciones mediante el intercambio de mensajes. Implementa el protocolo [AMQP][amqp] (Advanced Message Queuing Protocol) y es utilizado para enviar, recibir y encolar mensajes entre productores (aplicaciones que envían mensajes) y consumidores (aplicaciones que reciben mensajes).

### Características de RabbitMQ:

- Intermediación de Mensajes: RabbitMQ actúa como intermediario para transmitir mensajes entre diferentes aplicaciones, sistemas o servicios.
- Colas de Mensajes: Los mensajes se encolan hasta que un consumidor los recibe y los procesa.
- Soporte para Múltiples Protocolos: Aunque implementa principalmente AMQP, también soporta otros protocolos como MQTT y STOMP.
- Alta Disponibilidad: RabbitMQ puede ser configurado en clústeres para alta disponibilidad y recuperación ante fallos.
- Persistencia de Mensajes: Permite la persistencia de mensajes en disco para asegurar que no se pierdan en caso de fallos.
- Enrutamiento Flexible: Ofrece varias formas de enrutamiento de mensajes, incluyendo intercambio directo, intercambio de temas y enrutamiento basado en encabezados.
- Seguridad: Soporta autenticación y autorización, así como cifrado para asegurar la transmisión de mensajes.
- Monitoreo y Gestión: Incluye herramientas de administración y monitoreo a través de una interfaz web, así como APIs para gestión programática.

### Casos de Uso:

- Sistemas Distribuidos: Para comunicación fiable entre microservicios.
- Colas de Trabajo: Distribuir tareas entre múltiples trabajadores.
- Intermediación de Mensajes en Tiempo Real: Transmitir datos en tiempo real entre aplicaciones, como en sistemas de chat o notificaciones.
- Procesamiento de Eventos: Gestionar y procesar flujos de eventos en aplicaciones.

### Ejemplo de Funcionamiento:

1. Productor Envía Mensajes: Una aplicación (productor) envía un mensaje a una cola en RabbitMQ.
2. Cola de Mensajes: RabbitMQ almacena el mensaje en una cola.
3. Consumidor Recibe Mensajes: Otra aplicación (consumidor) recibe y procesa los mensajes de la cola.

RabbitMQ es ampliamente utilizado en arquitecturas modernas de microservicios y en aplicaciones que requieren una comunicación asincrónica y fiable entre componentes.

# Configuración de Docker Compose

Para comenzar montaremos un servicio de Rabbit creando un archivo de docker-compose con el siguiente codigo

```yml
version: "3.2"
services:
  rabbitmq:
    image: rabbitmq:3-management-alpine
    container_name: 'rabbitmq'
    ports:
        - 5672:5672
        - 15672:15672
    volumes:
        - F:\Develop\volumes/data/:/var/lib/rabbitmq/
        - F:\Develop\volumes/log/:/var/log/rabbitmq
    networks:
        - rabbitmq_go_net

networks:
  rabbitmq_go_net:
    driver: bridge
```
Acá una breve descripción del archivo
- **image**: indica la imagen que se usara, en usamos la implementación con los binarios de Alpine con el pluggin de administración, la implementación de alpine es muy liviana y segura.
- **container_name**: este es el nombre que le damos a nuestro contenedor.
- **ports**: este es el listado de puertos y sus mapeos con los puertos de la maquina para la comunicación 
- **volumes**: estos son los volumenes que montaremos para persistir la data en caso de que querramos usarlo nuevamente, sin embargo, no es obligatorio usarlo y en caso de que lo haga debe cambiar las rutas a unas que concuerden con su maquina.
- **networks**: En este apartado especificamos el nombre que le damos a la red creada para el contenedor

Para validar que funciona correctamente usamos el siguiente comando
```bash
docker-compose up
```
Esto bajará la imagen y le aplicará las configuraciones necesarias para correr en la maquina una vez culmine se puede acceder a la interfaz de [administación](http://localhost:15672)

use las credenciales guest como usuario y contraseña

# Go productor

Ahora crearemos otra carpeta para mantener el codigo separado con el nombre que guste, en mi caso usare go-rabbit y creare el archivo go.mod para realizar la instalación de las librerias para el proyecto.
```bash
mkdir go-rabbit
cd go-rabbit
go mod init rabbitgo
go get github.com/streadway/amqp
```
Una vez instalada la libreria creare el archivo `productor.go`, el cual tiene el siguiente codigo
```go
package main

import (
	"log"

	"github.com/streadway/amqp"
)

// Aca manejamos la forma en la que los errores son mostrados por consola.
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
	body := "Golang es el futuro!"
	err = ch.Publish(
		"",     // intercambio
		q.Name, // llave de enrutamiento
		false,  // obligatorio
		false,  // inmediato
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	// Si hay algun error publicando el mensaje se mostrara en la consola.
	failOnError(err, "Falla publicando el mensaje")
	log.Printf(" [x] Se ha enviado el mensaje: %s \n", body)
}

```

Podemos realizar la prueba del codigo corriendo `go run productor.go` se va a ver lo siguiente en consola.
```text
2024/05/26 12:40:56  [x] Se ha enviado el mensaje: Golang es el futuro! 
```


# Go consumidor

Ahora crearemos el archivo `consumidor.go` el cual tendrá el siguiente codigo
```go
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
```

Ahora que se ha creado un consumidor lo iniciamos con el comando ` go run consumidor.go` y vamos a ver lo siguiente en la consola

```text
2024/05/26 12:45:14  [*] Waiting for messages. To exit press CTRL+C
2024/05/26 12:45:14 Received a message: Golang es el futuro!
```
La linea que contiene `<- forever` indica que el canal permanecera abierto para siempre escuchando mensajes

# Go productor dinamico

Ya se ha realizado la prueba exitosa de la funcionalida del productor y el consumidor realizaremos un cambio en el codigo del productor para que nos pida el mensaje en cada ejecución agregando lo siguiente justo despues de `func main() {`
```go
// Tomamos el mensaje desde la terminal
reader := bufio.NewReader(os.Stdin)
fmt.Println("¿Qué mensaje quieres enviar?")
mPayload, _ := reader.ReadString('\n')
```
reemplazamos el body por mPayload dejandolo así 

```go
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
```







[rabbitmq]:https://www.rabbitmq.com
[amqp]:https://es.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol