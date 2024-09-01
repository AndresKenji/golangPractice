package main

import (
 "fmt"
 "sync"
 "time"
)

// Patrón Publicador-Suscriptor El patrón Publicador-Suscriptor permite que los mensajes se publiquen a múltiples suscriptores.
// Este patrón es útil en sistemas donde diferentes servicios necesitan reaccionar de manera independiente a ciertos eventos o tipos de mensajes.

// PubSub Esta estructura contiene un mapa de canales, channels, que utiliza los nombres de los temas (topics) 
// como claves y una lista de canales de cadenas ([]chan string) como valores.
// También tiene un Mutex (mu) para manejar la concurrencia, asegurando que las operaciones sobre los canales sean seguras
// en un entorno con múltiples goroutines.
type PubSub struct {
 mu       sync.Mutex                // Mutex to ensure safe concurrent access to channels map
 channels map[string][]chan string // Map of channels keyed by topic
}

// NewPubSub Esta función crea y devuelve una nueva instancia de PubSub inicializando el mapa de channels.
func NewPubSub() *PubSub {
 return &PubSub{
  channels: make(map[string][]chan string), // Initialize the map to store channels
 }
}

// Subscribe Este método permite a un suscriptor registrarse en un tema específico.
// Crea un nuevo canal, lo añade a la lista de canales asociados con ese tema,
// y luego devuelve el canal al suscriptor para que pueda recibir mensajes.
// Se usa un bloqueo (mu.Lock) para garantizar que la adición de canales sea segura.
func (ps *PubSub) Subscribe(topic string) <-chan string {
 ch := make(chan string) // Create a new channel for the subscriber
 ps.mu.Lock()
 ps.channels[topic] = append(ps.channels[topic], ch) // Add the channel to the list for the topic
 ps.mu.Unlock()
 return ch // Return the channel to the subscriber
}

// Publish Este método publica un mensaje en un tema específico.
// Para cada canal asociado con el tema, envía el mensaje a través del canal.
// De nuevo, se utiliza un bloqueo para asegurar que las operaciones concurrentes en el mapa de channels sean seguras
func (ps *PubSub) Publish(topic, msg string) {
 ps.mu.Lock()
 for _, ch := range ps.channels[topic] {
  ch <- msg // Send the message to each channel for the topic
 }
 ps.mu.Unlock()
}

// Close Este método cierra todos los canales asociados con un tema determinado.
// Esto es útil cuando ya no se publicarán más mensajes en ese tema,
// permitiendo a los suscriptores salir de sus bucles de recepción de mensajes.
func (ps *PubSub) Close(topic string) {
 ps.mu.Lock()
 for _, ch := range ps.channels[topic] {
  close(ch) // Close each channel to signal that no more messages will be sent
 }
 ps.mu.Unlock()
}

func main() {
 ps := NewPubSub() // Crea una nueva instancia de PubSub.

 // Dos suscriptores se suscriben al tema "news".
 subscriber1 := ps.Subscribe("news")
 subscriber2 := ps.Subscribe("news")

 var wg sync.WaitGroup
 wg.Add(2) 
 // Se crean dos goroutines, cada una asociada con un suscriptor. 
 // Estas goroutines escuchan en sus respectivos canales y procesan los mensajes que reciben.

 // Goroutine for subscriber1
 go func() {
  defer wg.Done() // Signal when this goroutine is done
  for msg := range subscriber1 {
   fmt.Println("Subscriber 1 received:", msg) // Print received messages
  }
 }()

 // Goroutine for subscriber2
 go func() {
  defer wg.Done() // Signal when this goroutine is done
  for msg := range subscriber2 {
   fmt.Println("Subscriber 2 received:", msg) // Print received messages
  }
 }()

 // Publish Se publican dos mensajes en el tema "news", que ambos suscriptores recibirán.
 ps.Publish("news", "Breaking News!")
 ps.Publish("news", "Another News!")

 time.Sleep(time.Second) // Wait to ensure all messages are processed
 ps.Close("news") // Después de una pausa (1 segundo), el tema "news" se cierra, lo que provoca que los suscriptores terminen sus operaciones.
 wg.Wait() // Se espera que ambas goroutines terminen usando WaitGroup.