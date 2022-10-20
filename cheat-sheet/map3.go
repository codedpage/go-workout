package main

import "fmt"

func main() {

	type Vertex struct {
		X, Y float32
	}

	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)

}
