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

	fmt.Println(quote.Go())
}

func shivam() {
	fmt.Println("shivam!")
}
