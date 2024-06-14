package main

import (
	"log"
)

type Message struct {
	Name string `json:"name"`
	Msg string `json:"msg"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func logOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

