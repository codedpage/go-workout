package main

import "fmt"

func main() {
	done := make(chan int)
	fmt.Println("Hi")

	for i := 1; i <= 5; i++ {
		go hello(i, done)
	}

	for i := 1; i <= 5; i++ {

		fmt.Println("Finish ----", <-done)
	}

	fmt.Println("main")
}

func hello(i int, done chan int) {

	fmt.Println("Start ----", i)
	done <- i
}

//https://go.dev/play/p/GBzlRbM6Plf
