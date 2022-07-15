package main

import "fmt"

func main() {

	var n int
	var name string

	for {

		fmt.Printf("Enter roll no: ")
		fmt.Scanf("%d", &n)

		fmt.Printf("Enter  Name: ")
		fmt.Scanf("%s", &name)

		fmt.Println("Roll =", n, "\nName = ", name, "\n+++")
	}
}

//https://go.dev/play/p/OmN_iT__FbJ
