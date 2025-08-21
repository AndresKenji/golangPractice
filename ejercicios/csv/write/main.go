package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"name","age","gender"}
	data := [][]string{
		{"Alice","25","Female"},
		{"Oscar","37","Male"},
		{"Leidy","40","Female"},
	}
	writer.Write(headers)
	for _, row := range data {
		writer.Write(row)
	}
}