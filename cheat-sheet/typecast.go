package main
import "fmt"
func main() {	
	
var i int = 10
var f float64 = float64(i)
var u uint = uint(f)

// alternative syntax
i1 := 42
f1 := float64(i)
u1 := uint(f)


fmt.Print(i); fmt.Print("-------"); fmt.Println(i1);
fmt.Print(f); fmt.Print("-------"); fmt.Println(f1);
fmt.Print(u); fmt.Print("-------"); fmt.Println(u1);
}
