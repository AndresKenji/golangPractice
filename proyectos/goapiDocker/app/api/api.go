package api

import (
	"log"
	"net/http"
)

type APIserver struct {
	addr string
}

type Middleware func(next http.Handler) http.HandlerFunc

func NewApiServer(addr string) *APIserver {
	return &APIserver{addr: addr}
}

func (s *APIserver) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", index)
	mux.HandleFunc("POST /slack", SendSlackMsgHandler)
	mux.HandleFunc("POST /teams", SendTeamsMsgHandler)

	// listado de middlewares
	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(mux),
	}

	log.Printf("Server running on %s", s.addr)

	return server.ListenAndServe()

}

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request method: %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
