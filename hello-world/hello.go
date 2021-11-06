package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Shivam Singhal")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Shivam", "nileso", "simple"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	fmt.Println(quote.Go())
}

func shivam() {
	fmt.Println("shivam!")
}
