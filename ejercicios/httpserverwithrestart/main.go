package main

import (
	"context"
	"log"
	"net/http"
)

type server struct {
	port        string
	Mux         *http.ServeMux
	Context     context.Context
	restartChan chan bool
}

func newServer(context context.Context, restar chan bool) *server {
	log.Println("Creating server")
	server := &server{}
	server.Context = context
	server.restartChan = restar
	server.port = "8080"
	return server
}

func (s *server) setRoutes() {
	log.Println("Setting server routes")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from api!"))
	})
	mux.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		s.restartChan <- true
		w.Write([]byte("Restart signal sent"))
	})

	s.Mux = mux
}

func (s *server) startServer() error {
	srv := &http.Server{
		Addr:    ":" + s.port,
		Handler: s.Mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	log.Printf("Server started on port %s\n", s.port)

	<-s.Context.Done()

	log.Println("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		return err
	}

	log.Println("Server stopped gracefully.")
	return nil
}
func main() {

	for {
		checkForRestart := make(chan bool)
		ctx, cancel := context.WithCancel(context.Background())
		srv := newServer(ctx, checkForRestart)
		srv.setRoutes()

		go func() {
			log.Println("Iniciando server")
			if err := srv.startServer(); err != nil {
				log.Panic(err)
			}
		}()

		select {
		case <-checkForRestart:
			log.Println("restarting server...")
			cancel()
		case <-ctx.Done():
			log.Println("Context canceled, shutting down the server...")
			return
		}

	}
}
