package main

import (
	"example.com/methods/interfaces"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X + v.Y
}

func main() {
	myV := Vertex{12, 12}
	fmt.Println(myV.Abs())

	interfaces.MainFunc()
}
