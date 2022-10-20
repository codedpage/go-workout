package main
import "fmt"
import "strconv"
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


//string concatenation
    a := "10"
    b := "20"
    c := fmt.Sprintf("%s%s",a,b) 
    d,_ := strconv.ParseInt(c,10,64)
    fmt.Println(d*2)

//number concatenation
    x := 20
    y := 30
    z := fmt.Sprintf("%d%d",x,y) 
    e,_ := strconv.ParseInt(z,10,64)
    fmt.Println(e*2)



//#EOF
}