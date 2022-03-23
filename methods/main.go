package main

import (
	"example.com/methods/interfaces"
	"fmt"
	"time"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X + v.Y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X: %v \nY: %v", v.X, v.Y)
}

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}

type MyError struct {
	When time.Time
	What string
}

type Joker struct {
	smile bool
}

func (e *MyError) Error() string {
	return fmt.Sprintf("hi at %v, %s",
		e.When, e.What)
}

func (e *Joker) Error() string {
	return "ss"
}

func run() error {
	return &Joker{
		true,
	}
}

func Foo(errTrue bool) (error, string) {
	if errTrue {
		return &MyError{
			time.Now().Add(8 * time.Hour),
			"something didnt work out",
		}, "nil"
	} else {
		return nil, "success"
	}
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	err, value := Foo(true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value: ", value)
	}

	myV := Vertex{12, 12}
	fmt.Println(myV.Abs())

	fmt.Println(myV)
	interfaces.MainFunc()
}
