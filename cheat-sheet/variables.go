package main
import "fmt"
func main() {

//Bool
    myBool := true
    if myBool {
        fmt.Println("T")
    } else {
        fmt.Println("F")
    }

//Int & Float
    myInt := 2
    myFloat := 1.5

    calInt := myInt * int(myFloat)
    fmt.Println(calInt)

    calFloat := float64(myInt) * myFloat
    fmt.Println(calFloat)

//#EOF
}