package main

import "fmt"

func main() {
	// switch statement
	
	var operatingSystem string = "linux"
	
	switch operatingSystem {
	case "darwin":
		fmt.Println("Mac OS Hipster")
	// cases break automatically, no fallthrough by default
	case "linux":
		fmt.Println("Linux Geek")
	default:
		// Windows, BSD, ...
		fmt.Println("Other")
	}
}
