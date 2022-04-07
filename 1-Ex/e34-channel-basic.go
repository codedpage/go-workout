package main

import "fmt"

func main() {
	var a chan int
	a = make(chan int)
	fmt.Printf("%T %v\n", a, a)
}

//https://go.dev/play/p/6qpQmdTeQ3c
