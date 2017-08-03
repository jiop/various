package main

import (
	"fmt"
	"time"
)

const MAX = 100

func test(c chan<- int) {
	for i := 0; i <= MAX; i++ {
		c <- i
		time.Sleep(10 * time.Millisecond)
	}
	close(c)
}

func test2(c chan<- int) {
	for i := 0; i <= MAX*10; i++ {
		c <- i
		time.Sleep(1 * time.Millisecond)
	}
	close(c)
}

func test3(c chan<- int) {
	c <- 10
	close(c)
}

func main() {
	c := make(chan int)
	go test(c)
	go test2(c)
	for v := range c {
		fmt.Println("Received ", v)
	}

	c3 := make(chan int)
	go test3(c3)
	for v := range c3 {
		fmt.Println("Received 3", v)
	}
}
