package main
import "fmt"
func main() {
	//var a []int               			// a slice – like an array, but length is unspecified
	//var a = []int{1, 2, 3, 4} 			// declare and initialize a slice
	//a := []int{1, 2, 3, 4}    			// shorthand
		
	a:= []string{0: "a", 2: "c", 1: "b"} 		// ["a", "b", "c"]

	fmt.Println("Hello Go")
	for i, e := range a {
		fmt.Println(i)
		fmt.Println("------------")
		fmt.Println(e)
	}
//////////////////////////////////////////////////////////////////	
	
	s:= []string{"a","b","c","d","e","f","g","h","i","j"} 
	var b = s[2:5]			//c,d,e        
	//var b =  s[:2]		//a,
	//var b =  s[2:]		//c,d,e...,i,j
	
fmt.Println("=============")
	
for i, e := range b {
		fmt.Print(i)
		
		fmt.Println(" - ", e)
	}	
}
