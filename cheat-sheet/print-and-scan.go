package main
import "fmt"
func main() {
fmt.Println("Hello Go")

fmt.Print("Enter text: ")
var input string
fmt.Scanln(&input)
fmt.Print(input)

fmt.Print("Enter no: ")
var no int
fmt.Scanln(&no)
fmt.Print(no)


}