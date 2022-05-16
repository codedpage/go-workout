package main

import "fmt"

func main() {
	c := make(chan int)
	go prod(c)

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func prod(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	//close(c)
}

//https://go.dev/play/p/Hgjay901R3t
