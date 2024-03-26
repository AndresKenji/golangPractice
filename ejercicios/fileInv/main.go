package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	var route string
	var outFile string

	flag.StringVar(&route, "root", "/", "Ruta a examinar")
	flag.StringVar(&outFile, "outfile", "data", "Nombre del archivo de salida")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Uso del programa:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	//root := "F:/Kenji"
	files := make([]fileInfo, 0, 200)
	saveDirFiles(&files, route)

	fmt.Println("El total de archivos obtenidos es de", len(files))

	// Crear un nuevo archivo CSV para escribir
	file, err := os.Create(outFile + ".csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Crear un escritor CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir el encabezado del CSV
	header := []string{"Full_name", "Name", "size", "extension"}
	if err := writer.Write(header); err != nil {
		log.Panic(err)
	}

	// Escribir los datos de cada estructura en el archivo CSV
	for _, fileInfo := range files {
		record := []string{
			fileInfo.FullName,
			fileInfo.Name,
			strconv.FormatInt(fileInfo.Size, 10),
			fileInfo.Extension,
		}
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}

}

func saveDirFiles(files *[]fileInfo, directory string) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.IsDir() {
			saveDirFiles(files, filepath.Join(directory, e.Name()))
		} else {
			inf, _ := e.Info()
			var newFile fileInfo
			newFile.FullName = filepath.Join(directory, e.Name())
			newFile.Name = e.Name()
			newFile.Size = inf.Size()
			newFile.Extension = filepath.Ext(filepath.Join(directory, e.Name()))

			*files = append(*files, newFile)
		}
	}
}

type fileInfo struct {
	FullName  string
	Name      string
	Size      int64
	Extension string
}
