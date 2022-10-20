package main

import "fmt"

func main() {
	a := 10; 
	p := &a

fmt.Println(p)  	 	 
fmt.Println(*&a)  

a1 := "abc"
p1 := &a1

fmt.Printf("%p", p1)
fmt.Printf("%s", *p1)


	 
}
