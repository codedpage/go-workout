package main

import "fmt"

type Employee struct {
	name string
	age  int
}

func (e Employee) tch(s string) {
	e.name = s
	fmt.Println(e)
}

func (e *Employee) pch(s string) {
	e.name = s
	fmt.Println(e)
}

func main0() {
	e := Employee{"aa", 20}
	fmt.Println(e)
	e.tch("bb")
	fmt.Println(e)
	(&e).pch("cc")
	fmt.Println(e)
}

func main() {
	v := Employee{"aa", 20}
	r := &v
	fmt.Println(v)
	fmt.Println("---\n")

	//v-v
	v.tch("bb")
	fmt.Println(v)
	fmt.Println("vv---tc \n")

	//r-r
	r.pch("cc")
	fmt.Println(v)
	fmt.Println("rr---pc\n")

	//r-v
	r.tch("dd")
	fmt.Println(v)
	fmt.Println("rv---tc \n")

	//v-r
	v.pch("ee")
	fmt.Println(v)
	fmt.Println("vr---pc \n")
}

//https://go.dev/play/p/920r3fAohvl
