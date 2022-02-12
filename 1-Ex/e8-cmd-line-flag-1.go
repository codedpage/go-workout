package main

import (
	"flag"
	"fmt"
)

func main() {

	a := flag.String("word", "foo", "a string")

	var b string
	flag.StringVar(&b, "text", "bar", "a string")

	flag.Parse()

	c := &b
	fmt.Println("a - word : ", *a, a)
	fmt.Println("b - text : ", b)
	fmt.Println("c - text : ", *c, c)

	fmt.Printf("\n%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)

}
