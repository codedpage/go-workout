// There's only `for`. No `while`, no `until`
package main
import "fmt"
func main() {
fmt.Println("For loop : ex-1")
for i := 1; i <= 3; i++ {
	fmt.Println(i)
}
fmt.Println("For loop : ex-1")
j := 1
for ; j <= 3; { 
	if (j > 10){
		return;
	  }
	fmt.Println(j)  
	j++
}
fmt.Println("while loop : ex-1")
k := 1
for k <= 3 { 
	if (k > 10){
		return;
	  }
	fmt.Println(k)  
	k++
}
fmt.Println("while loop : ex-2")
n := 1
for  { 
	if (n > 3){
		return;
	  }
	fmt.Println(n)  
	n++
}
}


