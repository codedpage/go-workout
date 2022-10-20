package main

import "fmt"

func main() {

	// declare and initialize
	//var a = [2]int{1, 2}
	//a := [2]int{1, 3} //shorthand
	a := [...]int{1, 4} // elipsis -> Compiler figures out array length

	fmt.Println("Hello Go")
	for i, e := range a {
		fmt.Print(i)
		fmt.Print("------------")
		fmt.Println(e)
	}
}