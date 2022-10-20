package main

import "fmt"

func main() {
	// Basic one
	x := -3
	if x > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}

	fmt.Println("-------------")
	// You can put one statement before the condition
	var b, c int = 41, 31
	if a := b + c; a < 42 {
		fmt.Println(a)
	} else {
		fmt.Println(a - 12)
	}

	fmt.Println("-------------")
	// Type assertion inside if
	var val interface{}
	val = 20
	if str, ok := val.(int); ok {
		fmt.Println(str)
	}

	fmt.Println("-------------")
	str1, ok1 := val.(int)
	fmt.Println(str1)
	fmt.Println(ok1)
}
