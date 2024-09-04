package server

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
)

type Server struct {
	SubcriberMessageBuffer int
	Mux                    http.ServeMux
	SubscriberMutex        sync.Mutex
	Subscribers            map[*subscriber]struct{}
}

type subscriber struct {
	msgs chan []byte
}

func NewServer() *Server {
	s := &Server{
		SubcriberMessageBuffer: 10,
		Subscribers:            make(map[*subscriber]struct{}),
	}

	s.Mux.Handle("/", http.FileServer(http.Dir("./htmx")))
	s.Mux.HandleFunc("/ws", s.subscribeHandler)
	return s
}

func (s *Server) AddSubscriber(subscriber *subscriber) {
	s.SubscriberMutex.Lock()
	s.Subscribers[subscriber] = struct{}{}
	s.SubscriberMutex.Unlock()
	log.Println("Added subscriber", subscriber)
}

func (s *Server) RemoveSubscriber(subscriber *subscriber) {
	s.SubscriberMutex.Lock()
	delete(s.Subscribers, subscriber)
	s.SubscriberMutex.Unlock()
	close(subscriber.msgs)
	log.Println("Removed subscriber", subscriber)
}

func (s *Server) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.Subscribe(r.Context(), w, r)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (s *Server) Subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}
	defer c.CloseNow()

	subscriber := &subscriber{
		msgs: make(chan []byte, s.SubcriberMessageBuffer),
	}
	s.AddSubscriber(subscriber)
	defer s.RemoveSubscriber(subscriber)

	ctx = c.CloseRead(ctx)
	for {
		select {
		case msg := <-subscriber.msgs:
			writeCtx, cancel := context.WithTimeout(ctx, time.Second)
			err := c.Write(writeCtx, websocket.MessageText, msg)
			cancel()
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *Server) Broadcast(msg []byte) {
	s.SubscriberMutex.Lock()
	defer s.SubscriberMutex.Unlock()
	for subscriber := range s.Subscribers {
		select {
		case subscriber.msgs <- msg:
		default:
			// If the subscriber's message buffer is full, skip the message.
			log.Println("Skipping subscriber due to full buffer")
		}
	}
}