package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Ross", "Ted", "Mike"}
	message, err := greetings.Greets(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
