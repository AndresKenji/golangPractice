package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var startSize int
	var endSize int
	years := 0

	
	flag.IntVar(&startSize,"startSize",100, "Tamaño inicial")
	flag.IntVar(&endSize,"endSize",300, "Tamaño inicial")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Uso del programa:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Con una población inicial es de %d para obtener una población final de %d seran necesarios:",startSize, endSize)
	for startSize < endSize  {
		startSize = (startSize / 3) - (startSize / 4) + startSize
		years ++
	}
	fmt.Printf("%d años",years)

}

