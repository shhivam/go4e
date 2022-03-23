package channels

import (
	"fmt"
	"time"
)

func testChannel(ch chan int) {
	ch <- 12
}

func ChannelMain() {
	fmt.Println("Hello channels main")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c2 <- "world"
	}()

	go func() {
		time.Sleep(time.Second)
		c1 <- "hello"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Ping: ", msg1)
		case msg2 := <-c2:
			fmt.Println("Ping:", msg2)
		}
	}
	ch := make(chan int)
	go testChannel(ch)
	fmt.Println("something: ", <-ch)
	time.Sleep(2 * time.Second)
	fmt.Println("something: ", <-ch)
	fmt.Println("Channle Main ended")
}
