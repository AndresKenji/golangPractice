package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T){
	for i:=0; i < 3; i++ {
		round := PlayRound(i)
		switch i {
		case 0:
			fmt.Println("El jugador eligio prieda")
		case 1:
			fmt.Println("El jugador eligio papel")
		case 2:
			fmt.Println("El jugador eligio tijera")
		}
		fmt.Printf("Message: %s\n",round.Message)
		fmt.Printf("Computer choice: %s\n",round.ComputerChoice)
		fmt.Printf("Round Result: %s\n",round.RoundResult)
		fmt.Printf("Computer choice int: %d\n",round.ComputerChoiceInt)
		fmt.Printf("Computer score: %s\n",round.ComputerScore)
		fmt.Printf("PlayerScore: %s\n",round.PlayerScore)
		fmt.Println("------------------------------------------------------")
	}

}