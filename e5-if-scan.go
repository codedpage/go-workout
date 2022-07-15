package main

import "fmt"

func main() {

	var n int
	fmt.Printf("Enter the no: ")

	fmt.Scanf("%d", &n)
	fmt.Scan(&n)
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

//https://go.dev/play/p/OsIhos6wQwR