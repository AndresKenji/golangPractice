package main

import (
	"fmt"
	"net/http"
)

type Holaservicio struct{}

func (hs *Holaservicio) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "Text/html")

	documento := fmt.Sprintf(`
    <h1>Bienvenido!</h1>
    <p>Ruta de acceso: %s</p>
    `, req.URL.Path)

	rw.Write([]byte(documento))
}

func main() {
	panic(http.ListenAndServe(":8000",&Holaservicio{}))
}