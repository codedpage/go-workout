package main

import (
	. "fmt"
)

//char occurance
func main() {
	str := "hello"

	output := make(map[string]int)

	for k, v := range str {
		Println(k, v)
		output[string(v)]++
	}
	Println(output)
}

//https://go.dev/play/p/6vrT3gB2pcR
