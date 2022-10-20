package main
import "fmt"
func main() {
	// assign a function to a name
	add := func(a, b int) int {
		return a + b
	}
	
	// use the name to call the function
	fmt.Println(add(3, 4))
	fmt.Println("----------------");
	
	var fn, e = outer()
	fmt.Println(fn())
	fmt.Println(e)
	fmt.Println("----------------");	

	var fn2 = scope()
	fmt.Println(fn2())
	 
}

func scope() func() int {
	outer_var := 20
	foo := func() int { return outer_var }
	return foo
}

// Closures: don't mutate outer vars, instead redefine them!
func outer() (func() int, int) {
	outer_var := 2 // NOTE outer_var is outside inner's scope
	inner := func() int {
		outer_var += 99  // attempt to mutate outer_var
		return outer_var // => 101 (but outer_var is a newly redefined
		// variable visible only inside inner)
	}
	return inner, outer_var // => 101, 2 (still!)
}
