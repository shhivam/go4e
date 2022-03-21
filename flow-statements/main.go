package main

import (
	"fmt"
	"strings"
)

func fibonacci() func() int {
	last, lastSecond := 0, 1
	return func() int {
		nextFib := last + lastSecond
		lastSecond = last
		last = nextFib
		return nextFib
	}
}

func fiboMain() {
	// Practicing closures
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// Practicing maps
func WordCount(s string) map[string]int {
	splitWords := strings.Split(s, " ")
	var frequencies map[string]int = make(map[string]int)

	for _, wordStr := range splitWords {
		if count, ok := frequencies[wordStr]; ok {
			frequencies[wordStr] = count + 1
		} else {
			frequencies[wordStr] = 1
		}
	}

	return frequencies
}

func main() {
	fmt.Println(WordCount("shivam is at at nilenso"))
	fiboMain()

	switch 13 {
	case 12:
		fmt.Println("hola world")
	case 13:
		fmt.Println("no hola world")
	case 19:
		fmt.Println("no nohola world")
	default:
		fmt.Println("default case")
	}

}
