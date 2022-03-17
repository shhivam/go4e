package main

import (
	"fmt"
)

func Average(lst []float64) float64 {
	if len(lst) == 0 {
		return 0
	}
	total := 0.0
	for _, value := range lst {
		total += value
	}
	fmt.Println("my tot:", total)
	return total / float64(len(lst))
}

func main() {
	myAvg := Average([]float64{})

	fmt.Println("Something", myAvg)
}
