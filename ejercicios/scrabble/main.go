package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
)


func main() {
	var palabra1 string
	var palabra2 string

	flag.StringVar(&palabra1, "palabra1", "nada","Palabra numero uno")
	flag.StringVar(&palabra2, "palabra2", "nada","Palabra numero dos")
	flag.Usage = func (){
		fmt.Fprintf(os.Stderr, "Scrabble va a evaluar el valor de las dos palabras que se ingresen y determinara la ganadora")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NFlag() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	score1 := ComputeScore(palabra1)
	score2 := ComputeScore(palabra2)
	fmt.Printf("La palabra numero 1 es %s y tiene un valor de %d \n",palabra1, score1)
	fmt.Printf("La palabra numero 2 es %s y tiene un valor de %d \n",palabra2, score2)
	if score1 > score2 { fmt.Printf("La palabra %s es la ganadora \n",palabra1)}
	if score2 > score1 { fmt.Printf("La palabra %s es la ganadora \n",palabra2)}
	if score1 == score2 { fmt.Println("Empate")}

}

func ComputeScore(word string) int {

	minusWord := strings.ToLower(word)
	wordPoints := map[string]int{
		"a": 1, "b": 3, "c": 3, "d": 2, "e": 1, "f": 4, "g": 2, "h": 4, "i": 1, "j": 8, "k": 5, "l": 1, "m": 3, "n": 1, "o": 1,
		"p": 3, "q": 10, "r": 1, "s": 1, "t": 1, "u": 1, "v": 4, "w": 4, "x": 8, "y": 4, "z": 10}
	var score int = 0
	for _, l := range(minusWord){
		score += wordPoints[string(l)]
	}
	return score
}
