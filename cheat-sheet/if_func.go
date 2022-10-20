package main
import "fmt"
func main() {
output := func1(15); fmt.Println(output)
var a,b,c int = 3,5,7; fmt.Println(func2(a,b,c))
func3()
}
// Basic one
func func1(x int) int  {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

// You can put one statement before the condition
func func2(a,b,c int) int {
	if a := b + c; a < 42  {
		return a
	} else {
		return -a	 
	}
}

// Type assertion inside if
func func3()  {
	var val interface{}
	val = 9//"welcome"
	if str, ok := val.(string); ok {
		fmt.Println(str)
	} else {
		fmt.Println("Not a valid string")
	}
}