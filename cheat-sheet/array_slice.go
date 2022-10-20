package main

import (  
    "fmt"
    "reflect"
)

func main() {

	//[...] - array, [6] - array, [] - slice
	a := []int{10, 20, 30, 40, 50, 60}					
	//a := []string{"a", "b", "c", "d", "e", "f"}

	s := a[1:3]

	t := reflect.TypeOf(a)	
	k := t.Kind()
	v := reflect.ValueOf(s)		
	
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)
	fmt.Println("Value ", v)
	fmt.Printf("length of slice %d capacity %d", len(s), cap(s)) //length of is 2 and capacity is 3	
}
