package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// abrir el archivo con os
	file, err := os.Open("inventario.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Leer los datos en csv
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // permite un numero variable de campos
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%s", col)
		}
		fmt.Println()
	}

}
