package main
import "fmt"
func main() {

//No params
f1()	
	
//Multiple parameters diff type
f2("String,",10)

//Multiple parameters same type */
f3(10,20)

}//#main

func f1()  {
	fmt.Println("No params")	
}

func f2(p1 string,p2 int)  {
	fmt.Print("Multi params: ")	
	fmt.Println(p1,p2)
}

func f3(p1,p2 int)  {
	fmt.Print("Multi params with same type: ")	
	fmt.Println(p1+p2)
}

