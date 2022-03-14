package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello work!")
}

func singleRoutine() {
	// Experiment with commenting out these print & sleep statements and see how the goroutine works
	go hello()
	fmt.Println("time: ", time.Second)
	fmt.Printf("\ntime type: %T\n", time.Second)
	time.Sleep(1 * time.Second)

}

func main() {
	singleRoutine()
	fmt.Println("main thread ended")
}
