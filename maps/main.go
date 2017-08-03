package main

import "fmt"

func genRand1(chn chan<- int) {
	for i := 0; i < 10; i++ {
		chn <- i
	}
	close(chn)
}

func genRand2(chn chan<- int) {
	for i := 10; i < 20; i++ {
		chn <- i
	}
	close(chn)
}

func main() {
	m := make(map[int]int)
	chn := make(chan int)
	chn2 := make(chan int)
	go genRand1(chn)
	go genRand2(chn2)
	idx := 0
	for v := range chn {
		m[idx] = v
		idx++
	}
	for v := range chn2 {
		m[idx] = v
		idx++
	}
	fmt.Printf("%v\n", m)

	//

	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}

	for person, salary := range personSalary {
		fmt.Printf("person: %s - salary: %d\n", person, salary)
	}
}
