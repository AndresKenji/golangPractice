package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gobocopy/internal/config"
	"gobocopy/internal/engine"
)

func main() {
	opts, err := config.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error de parametros: %v\n", err)
		printUsage()
		os.Exit(1)
	}

	report, err := engine.Run(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error de ejecucion: %v\n", err)
		printReport(report)
		os.Exit(2)
	}

	printReport(report)
	if report.FailedFiles > 0 && opts.FailOnError {
		os.Exit(3)
	}
}

func printUsage() {
	fmt.Println("Uso:")
	fmt.Println("  gobocopy -source <origen> -dest <destino> [opciones]")
	fmt.Println("")
	fmt.Println("Ejemplo:")
	fmt.Println("  gobocopy -source C:/datos -dest D:/backup -threads 16 -mirror -exclude-ext .tmp,.log -report-json ./report.json")
}

func printReport(r engine.Report) {
	payload, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Printf("Reporte: %+v\n", r)
		return
	}
	fmt.Println(string(payload))
}
