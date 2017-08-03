package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(name string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Hello, I'm %s %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Nanosecond)
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 50; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring, I'm leaving")
}
