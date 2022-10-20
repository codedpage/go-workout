package main
import "fmt"
func main() {
	c1 := f1()
	fmt.Print("F1:")
	fmt.Println(c1)

	var x, y = f2(10, 20)
	fmt.Print("F2:")
	fmt.Print(x)
	fmt.Println(y)


	var x1, y1 = f3(10, 21)
	fmt.Print("F3:")
	fmt.Print(x1)
	fmt.Println(y1)

} //#main

func f1() int {
	return 42
}

func f2(a int, b int) (int, string) {
	return a + b, " is an even no"
}

func f3(a int, b int) (n int, str string) {
	n = a+b
	str = " is an odd no"
	return 
}
