package main

import (  
    "fmt"
    "reflect"
)

func hello(done chan bool) {  
    fmt.Println("Hello world")
    done <- true
}

func main() {
    a := make(chan bool)
    go hello(a)
    <-a
    fmt.Println("main function")

	t := reflect.TypeOf(a)	
	k := t.Kind()
	v := reflect.ValueOf(a)		
	
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)
	fmt.Println("Value ", v)	 
}


/*
Hello world
main function
Type  chan bool
Kind  chan
Value  0x432080
*/