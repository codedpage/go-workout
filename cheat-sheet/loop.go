package main

//import "fmt"

func main() {

	// There's only `for`. No `while`, no `until`////////////////////
	for i := 1; i < 10; i++ {
		print(i)
	}		
	print("\n");
	
	// while loop ///////////////////////////////////////////////
	var j,k int =  2, 3; 
	for ; j < 10; { 		
		print(j)
		j++;
	}
	print("\n");
	
	// can omit semicolons if there's only a condition //////////////////////////
	for k < 10 { 
		print(k)
		k++;	
	}		
	print("\n");
	
	// can omit the condition ~ while (true)/////////////////////////////////////
	c:=1;
	for { 
		print(c);
		if (c == 10){
		break
		}
		c++;		
	}	 
}
