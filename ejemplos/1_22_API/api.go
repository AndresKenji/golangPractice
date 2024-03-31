package main

import (
	"log"
	"net/http"
)

type APIserver struct {
	addr string
}

func NewAPIserver(addr string) *APIserver {
	return &APIserver{addr: addr}
}

func (s *APIserver) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		UserId := r.PathValue("id")
		w.Write([]byte("User id is :"+ UserId))
	})

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1/", router))

	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		RequireAuthMiddleware,
	)
	
	server := http.Server{
		Addr: s.addr,	
		Handler: middlewareChain(router),
	}

	log.Printf("Server has started at %s", s.addr)

	return server.ListenAndServe()
}

// Middleware

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path %s", r.Method, r.URL.Path)
		next.ServeHTTP(w,r)
	}
}

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if the user is authenticated
		token := r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unautorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

type Middleware func(next http.Handler) http.HandlerFunc
func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i --{
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}


