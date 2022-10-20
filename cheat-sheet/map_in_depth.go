package main
import "fmt"
func main() {

//Array - Slice (numeric)
    myArray := [10]int{0,1,2,3,4,5}
    fmt.Println(myArray)
    fmt.Println(myArray[1:3])

    for k,v := range myArray[1:4] {
        fmt.Println(k, " => ", v)
    }

//Array - Slice (string)
    animal := []string{"cat","dog","lion"}
    fmt.Println(animal)

    fmt.Println(animal[2])

    for _,v := range animal {
        fmt.Println( "=>",v)
    }

//multi dimension array
    alpha := map[string]int{
        "a":1,
        "b":2,
        "c":3,
    }
    alpha ["d"] = 4

    for k, v := range alpha {
        fmt.Println(k, " => ", v)
    }
    //---------------
    alpha2 := map[int]string{
        1:"a",
        2:"b",
        3:"c",
    }
    alpha2[4] = "d"

    for k, v := range alpha2 {
        fmt.Println(k, " => ", v)
    }
    //---------------
    alpha3 := map[string]string{
        "a":"aaa",
        "b":"bbb",
        "c":"ccc",
    }
    alpha3["d"] = "ddd"

    for k, v := range alpha3 {
        fmt.Println(k, " => ", v)
    }

    delete(alpha3, "b")
    e, ok := alpha3["b"]
    if ok{
        fmt.Println("Available",e)
    } else {
        fmt.Println("Not available",e)
    }
    fmt.Println(alpha3)

//Ex---------
    var m map[string]int
    m = make(map[string]int)
    m["key"] = 45
    fmt.Println(m)

    n := map[string]int{}
    n["key"] = 45
    fmt.Println(n)

    k := map[string]int{"key" : 45}
    fmt.Println(k)

//Ex---------
    type Point struct {
        x,y float32
    }

    var m1 = map[string]Point{
        "p1": {2.0,4.0},
        "p2":    {6.5,9.5},
    }
    fmt.Println(m1)

//Ex---------
    var data2 = map[string]string{}
    data2["a"] = "aa"
    data2["b"] = "bb"
    data2["c"] = "cc"
    fmt.Println(data2)
 
    var data11 = map[string][]string{}
    data11["a"] = append(data11["a"], animal[1])
    data11["b"] = append(data11["b"], "x","y","z")
    data11["c"] = append(data11["c"], "x","y","z")
    fmt.Println(data11)

//Ex--------- 
    //var s = map[int]string{}
    s := map[int]string{
            1:"x",
            2:"y",
            3:"z",
            }
    s[4] = "w"
    var data12 = map[string]map[int]string{}
    data12["a"] = s
    data12["b"] = s
    data12["c"] = s
    fmt.Println(data12)
 

    var data13 = map[string]map[string]string{
    "a": map[string]string{"1": "x", "2" : "y", "3":"z"},
    "b": map[string]string{"1": "x", "2" : "y", "3":"z"},
    "c": map[string]string{"1": "x", "2" : "y", "3":"z"},
    "d": map[string]string{},
    }
    fmt.Println(data13)
 
    var data14 = map[string]map[int]string{
    "a": map[int]string{1: "x", 2 : "y", 3:"z"},
    "b": map[int]string{1: "x", 2 : "y", 3:"z"},
    "c": map[int]string{1: "x", 2 : "y", 3:"z"},
    }
    fmt.Println(data14)

//Ex---------
    var data3 = map[string]map[string]string{
        "a": map[string]string{},
        "b": map[string]string{},
        "c": map[string]string{},
    }
    data3["a"]["w"] = "x"
    data3["b"]["w"] = animal[1]
    data3["c"]["w"] = "y"
    fmt.Println(data3) 

    var data31 = map[string]map[string]string{}
    data31["a"] = map[string]string{}
    data31["b"] = make(map[string]string)
    data31["c"] = make(map[string]string)

    data31["a"]["w"] = "x"
    data31["b"]["w"] = animal[1] 
    data31["c"]["w"] = "y"
    fmt.Println(data31)

    for k,v := range data31 {
     //fmt.Println(k,v)
     fmt.Print(k, "-", v["w"],",")
    }

//#EOF
}