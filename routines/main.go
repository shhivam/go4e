package main

import (
	"example.com/routines/channels"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

func multipleRoutines() {
	fmt.Println("Multiple routines starting!")
	go numbers()
	go alphabets()
	time.Sleep(400 * time.Millisecond)
	fmt.Println("multi routines ending")
}

func hello() {
	fmt.Println("Hello work!")
}

func singleRoutine() {
	// Experiment with commenting out these print & sleep statements and see how the goroutine works
	go hello()
	fmt.Println("time: ", time.Second)
	fmt.Printf("\ntime types: %T\n", time.Second)
	time.Sleep(1 * time.Second)

}

func newtonSqrRoot(n float64, tolerance float64) float64 {
	var x = n
	var root float64 = 0
	for {
		root = 0.5 * (x + (n / x))
		if math.Abs(root-x) < tolerance {
			break
		}
		x = root
	}

	return root
}

func add(x int, y int) int {
	fmt.Println("Result: ", x+y)
	return 0
}

func main() {
	singleRoutine()
	fmt.Println("main thread starting")
	multipleRoutines()
	fmt.Println("	main thread ended")
	c := complex(12, 11)
	const hi = complex(12, 12)
	const (
		Big float64 = 1 << 100
	)
	fmt.Printf("%T", Big)
	fmt.Println(real(c))
	fmt.Println(imag(c))
	fmt.Println(real(hi))
	fmt.Println(imag(hi))
	fmt.Println(newtonSqrRoot(16, 0.001))
	resp, err := http.Get("http://127.0.0.1:5000/shivam")
	if err != nil {
		fmt.Println("Went wrong!", err)
	}

	fmt.Println("resp --> ", resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("parsing went wrong!")
	}
	fmt.Println("Parsed Body: ", string(body))
	fmt.Println("Start")
	defer fmt.Println("defered Middle order")
	defer add(34, 56)
	defer add(10, 10)

	channels.ChannelMain()
}
