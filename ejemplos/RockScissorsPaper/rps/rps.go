package rps

import (
	"math/rand"
	"strconv"

)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choiceInt"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"¡Bien hecho!",
	"¡Buen trabajo!",
	"¡Eres una maquina!",
	"¡Crack!",
}

var loseMessages = []string{
	"¡Ups!",
	"¿Quieres llorar?",
	"¡Eres una verguenza!",
	"Bueno no todos son ganadores",
}

var drawMessages = []string{
	"Buen empate, mal contrincante",
	"¿Acaso eres tan bueno como yo?",
	"¿Jugamos otra?",
	"Casi pierdes",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligio TIJERAS"
	}

	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	}else if playerValue == (computerValue+1)%3 {
		PlayerScore ++
		roundResult = "El jugador gana"
		message = winMessages[messageInt]

	}else {
		ComputerScore ++
		roundResult = "El jugador pierde"
		message = loseMessages[messageInt]
	}

	return Round{
		Message: message,
		ComputerChoice: computerChoice,
		RoundResult: roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore: strconv.Itoa(ComputerScore),
		PlayerScore: strconv.Itoa(PlayerScore),
	}
}