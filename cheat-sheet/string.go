package main
import "fmt"
func main() {
    name := "manojk"
    
    for _, v := range []byte(name) {         
        //fmt.Println(k, v)
        fmt.Printf("%c",v)

    }

    fmt.Println()

    for _, v1 := range []rune(name) {         
        //fmt.Println(k1, v1)
        defer fmt.Printf("%c",v1)
    }

    fmt.Println("\n")

    for k2, v2 := range (name) {         
       fmt.Println(k2, v2)
    }
	
/////////////////////////////////

	name := "manojk"
	printCharsAndBytes(name)
	
}	

 func printCharsAndBytes(s string) {
	for i, int32:= range s {
		fmt.Printf("%c starts at byte %d\n", int32, i)
	}
 }

