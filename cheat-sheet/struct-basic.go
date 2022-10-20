// You can edit this code!
// Click here and start typing.
package main

func main() {

	//There are no classes, only structs. Structs can have methods.
	// A struct is a type. It's also a collection of fields
	// Declaration
	type Vertex struct {
		X, Y int
	}
	// Creating
	//var v = Vertex{1, 2}
	var v = Vertex{X: 1, Y: 2} // Creates a struct by defining values
	// with keys

    print(v.X, v.Y);
    print("\n");
	// Accessing members
	v.X = 4
    print(v.X, v.Y);
}
