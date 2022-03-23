package main

import (
	"fmt"
	"sync"
)

var GFG = 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {

	m.Lock()
	GFG += 1
	m.Unlock()
	defer wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w, &m)
	}

	w.Wait()
	fmt.Println("Value of X:", GFG)
}
