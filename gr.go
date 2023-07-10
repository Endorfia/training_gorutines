package main

import "fmt"

func endo(a, b chan int) {
	for {
		n := <-a
		b <- n * n
	}
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go endo(a, b)

	for i := 0; i < 10; i++ {
		a <- i
		fmt.Println(<-b)
	}
}
