package api

import (
	"log"
	"net/http"
)

type APIserver struct {
	addr string
}

// Metodo "constructor" de la api
func NewApiServer(addr string) *APIserver { return &APIserver{addr: addr} }

// Run es el metodo usado para iniciar el server y en el que realmente se encuentra la logica del paquete http
func (s *APIserver) Run() error {
	// router se encarga de ejecutar la logica segun la ruta solicitada
	router := http.NewServeMux()
	// al agregar un handlefunc a un router estamos especificando el Metodo servidor/ruta y la función a ejecutar
	router.HandleFunc("GET /",
		index)
	router.HandleFunc("GET /cookie",
		CheckCookieHandler)
	router.HandleFunc("POST /login",
		LoginHandler)
	router.HandleFunc("GET /users",
		RequireLoginMiddleware(GetUsersHandler))
	router.HandleFunc("GET /users/{id}",
		RequireLoginMiddleware(GetUserHandler))
	router.HandleFunc("POST /users",
		PostUserHandler)
	router.HandleFunc("DELETE /users/{id}",
		RequireLoginMiddleware(DeleteUserHandler))
	router.HandleFunc("GET /tasks",
		RequireLoginMiddleware(GetTasksHandler))
	router.HandleFunc("GET /tasks/{id}",
		RequireLoginMiddleware(GetTaskHandler))
	router.HandleFunc("PATCH /tasks/{id}",
		RequireLoginMiddleware(UpdateTaskHandler))
	router.HandleFunc("POST /tasks",
		RequireLoginMiddleware(PostTasksHandler))
	router.HandleFunc("DELETE /tasks/{id}",
		RequireLoginMiddleware(DeleteTaskHandler))

	// listado de middlewares
	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
	)

	// El servidor http es el que estará encargado de escuchar las peticiones en la dirección url indicada y pasa estas peticiones al handler(router)
	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}
	log.Printf("Server running on %s", s.addr)
	return server.ListenAndServe()
}

// Middleware es una función que toma como parametro un handler de http y retorna una nueva función con alguna logica enriquesida o con ciertos criterios
type Middleware func(next http.Handler) http.HandlerFunc

// MiddlewareChain es una funcion que toma como parametro cualquier numero de Middlewares y retorna un unico midleware que ejecuta todas las logicas de las demas middlewares.
// El orden en el que se indiquen los middlewares es el mismo orden en el que van a ser ejecutados
func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}

// RequestLoggerMiddleware es un middleware que toma como argumento un handler http y ejecuta una logica previa a la ejecución del handler, en este caso imprime el metodo y la ruta solicitada
func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request method: %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

// RequireLoginMiddleware es un middleware que verifica que se tenga un header de autenticación valido y de ser asi permite la ejecución de un handler
func RequireLoginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing authorization header"))
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token"))
			return
		}
		next(w, r)
	}
}
