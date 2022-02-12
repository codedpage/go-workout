package main

import "fmt"

func main() {
	a := [3][2]string{{"aa", "bb"}, {"mm", "nn"}, {"xx", "yy"}}
	arrPrint(a)

}

func arrPrint(a [3][2]string) {

	for _, v := range a {
		fmt.Print(v)
		for _, v1 := range v {
			fmt.Print(v1, " ")
		}
		fmt.Println()
	}
}
