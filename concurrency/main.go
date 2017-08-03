package main

import (
	"fmt"
	"math/rand"
	"time"
)

const base = 100

func concurrentFunc() (results []int) {
	c := make(chan int)
	go func() {
		time.Sleep(time.Duration(rand.Intn(base)) * time.Millisecond)
		c <- 1
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(base)) * time.Millisecond)
		c <- 2
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(base)) * time.Millisecond)
		c <- 3
	}()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case <-timeout:
			fmt.Println("timed out")
			return
		case result := <-c:
			results = append(results, result)
		}
	}
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := concurrentFunc()
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
