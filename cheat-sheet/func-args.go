package main

import "fmt"

func main() {
	fmt.Println(adder(1, 2, 3)) // 6
	fmt.Println(adder(9, 9))    // 18
	nums := []int{10, 20, 30}
	fmt.Println(adder(nums...)) // 60
}

// Using ... before the type name of the last parameter indicates
// that it takes zero or more of those parameters.
// The function is invoked like any other function except we can
// pass as many arguments as we want.
func adder(args ...int) int {
	total := 0
	for _, v := range args { // Iterate over all args
		total += v
	}
	return total
}
