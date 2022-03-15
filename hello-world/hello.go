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

	// Using a function from another file from the same package i.e. `main`
	FunctionFromAnotherFile()

	fmt.Println(quote.Go())
}

// func main() {
// 	fmt.Println("..")
// }

func shivam() {
	fmt.Println("shivam!")
}

func init() {
	// fmt.Println("init func")
	// var out []int
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("index: ", i)
	// 	out = append(out, i)
	// }
	// fmt.Println("Array: ", out)
	// var s *int
	// var mine int = 40
	// fmt.Println("s = ", s)
	// fmt.Println("address s = ", &s)
	// fmt.Println("address mine = ", &mine)
	// fmt.Println("hmmmm", *&mine)
	// fmt.Println("assigning")
	// s = &mine
	// fmt.Println("s = ", s)

	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
	shivam()
}

func init() {
	fmt.Println("h")
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("\n%c", i)
	}
}
