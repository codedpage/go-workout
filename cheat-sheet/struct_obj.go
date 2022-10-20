package main
import (
    "fmt"
)   
type sc struct {
  name string
}
type point struct {
  x,y float32
}
func print(t *sc) {   
    fmt.Println(t) //t.name
}
func print1(p point) {   
    fmt.Println(p) //p.x
}
func print2(t sc) {   
    fmt.Println(t) //p.x
}

func  main() {
    obj:= new(sc)
    obj.name= "India"
    print(obj)

    var m1 =  point{2.0,4.0}
    print1(m1) 

    var m2 =  sc{"ram"}
    print2(m2) 
    
    obj1:= new(sc)
    obj1.name= "India"
    fmt.Println(obj1) 

}