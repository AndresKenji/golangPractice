// Package hola implementa maneras de saludar
package hola

import "fmt"

// Con nombre retorna un efusivo saludo al nombre pasado como argumento
func ConNombre(nombre string) string {
	return fmt.Sprintf("¡Hola, %s", nombre)
}
