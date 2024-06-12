# Rabbit MQ & Go

En este proyecto usaremos RabitMQ con docker compose y crearemos un productor y un consumidor usando go, para correr esto en su maquina asegurese de tener instalado lo siguiente:

- Docker
- Docker compose
- Go

<br>

## Rabbitmq

[RabbitMQ][rabbitmq] es un software de código abierto para el manejo de colas de mensajes, también conocido como message broker o middleware de mensajería. Facilita la comunicación entre aplicaciones al enviar, recibir y gestionar mensajes de manera eficiente y fiable. 

### Características de RabbitMQ

1. Protocolo AMQP: RabbitMQ implementa el protocolo AMQP (Advanced Message Queuing Protocol), que define la semántica para la mensajería, garantizando una comunicación segura y fiable entre los sistemas.

2. Colas de Mensajes: Permite la creación de colas donde los mensajes se almacenan hasta que son procesados por un consumidor.

3. Intercambios (Exchanges): Utiliza intercambios para dirigir los mensajes a las colas. Los intercambios pueden ser de diferentes tipos, como directos, fanout, tópicos y encabezados, dependiendo de cómo se quiera enrutar el mensaje.

4. Persistencia: Los mensajes pueden ser persistentes (guardados en el disco) para asegurar que no se pierdan en caso de fallos del sistema.

5. Alta Disponibilidad: Soporta la replicación de colas en múltiples nodos para asegurar la disponibilidad y la tolerancia a fallos.

6. Plugins y Extensiones: RabbitMQ tiene una arquitectura extensible que permite la incorporación de plugins para añadir funcionalidades adicionales.

7. Herramientas de Administración: Ofrece una interfaz de gestión web y una serie de herramientas de línea de comandos para la administración y monitoreo de mensajes, colas, intercambios y más.

### Entidades de Rabbit

[RabbitMQ][gotoiot] utiliza terminos que definen el comportamiento de cada entidad involucrada en el envio y la recepción de los mensajes.

- **Productor:** Todo programa que envie mensajes es un productor
- **Cola:** Es un buffer de memoria en el que se contiene a los mensajes que son enviados por los productores y seran consumidos por algun consumidor.
- **Consumidorr:** Todo programa que espera recibir mensajes de una cola es un consumidor y consumir significa recibir y procesar un mensaje.

![partes](./readme-img/rabbitparts.png)

Dentro del modelo AMQP 0-9-1, los productores envían mensajes hacia los exchanges indicando diferentes propiedades que son utilizadas por el broker para realizar el enrutamiento. Los exchanges distribuyen copias de los mensajes a las colas vinculadas utilizando diferentes criterios. Los criterios de asociación entre exchanges y colas se denominan binding, y cada tipo de exchange posee sus propias reglas para enrutar los mensajes hacia las colas. Finalmente, el broker envía los mensajes a los consumidores suscritos a las colas mediante una notificación push, o bien mediante un mecanismo de polling. Veamos en esta imagen este modelo simplificado.

#### Exchanges

Los exchanges son los enrutadores de los mensajes a las colas, su objetivo es enrutar los mensajes a las colas que cumplan con los criterios si hay alguna. 

Los atributos mas importantes de un Exchange son:
- Name
- Durable (determina si sobrevive luego de un reinicio del broker)
- Auto-delete (Se elimina cuando la ultima cola asociada a el se libera)

El algoritmo de enrutamiento cambia según el tipo de exchange que se utilice los cuales son:

- **Default:** Es el tipo predefinido cada cola  que se crea se vincula automáticamente con una routing_key que es el mismo nombre de la cola; Ejemplo si se declara una cola con el nombre de syslogs, el broker la vincula al default exchange utilizando syslogs como routing_key, cuando un mensaje es publicado en el default exchange con la routing_key syslogs se enruta a la cola syslogs.

- **Direct:** Entrega mensajes a las colas según la routing_key con la que se publican los mensajes. Las colas se vinculan al exchange con una routing_key y al llegar un mensaje nuevo al exchange con la misma routing_key es enviado a la cola.

	![directo](./readme-img/directExchange.png)

- **Fanout:** Envia los mensajes a todas las colas que se encuentran vinculadas a el sin tomar en cuenta la routing_key del mensaje ni del binding.

	![directo](./readme-img/fanoutExchange.png)

- **Topic:** Enruta los mensajes a las colas que según la concidencia entre la routing_key del mensaje y la routing_key con la que se realiza el binding entre el exchange y las colas.<br>
Los mensajes enviados a un topic exchange tienen una routing_key que son una lista de palabras delimitadas por puntos. Se puede implementar cualquier combinación de palabras siempre y cuando no superen los 255 bytes. Algunos ejemplos de routing_key válidas son "building1.room1.sensor3.temperature" o "chat5.user3.alert".<br>
Para que los mensajes publicados en un topic exchange puedan llegar a una cola, la routing_key utilizada en el binding que los vincula debe tener el mismo formato que la de publicación del mensaje. Así mismo, hay caracteres especiales que pueden dar flexibilidad a este esquema. El caracter asterisco (*) puede sustituir exactamente una palabra, mientras que el caracter numeral (#) puede sustituir varias palabras.

	![directo](./readme-img/topicExchange.png)

- **Header:** Un header exchange está diseñado para enrutar mensajes a través de varios atributos que se expresan en el encabezado del mensaje en vez de como routing_key. Un mensaje se considera coincidente si el valor del header es igual al valor especificado en el binding entre la cola y el exchange. <br>
Es posible vincular una cola a un header exchange usando más de un atributo en el header, lo que permite implementar mecanismos más complejos de enrutamiento de los que vimos hasta ahora. <br>

	![directo](./readme-img/headerExchange.png)



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
[gotoiot]:https://www.gotoiot.com/pages/articles/rabbitmq_intro/index.html#:~:text=Los%20exchanges%20AMQP%200-9,no%20implementar%20explícitamente%20ninguno%20estándar.
[rabbitmq-tutorial]:https://www.rabbitmq.com/tutorials/amqp-concepts