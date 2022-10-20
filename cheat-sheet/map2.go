package main
import "fmt"
func main() {

    var data1 = map[string]string{}
    data1["a"] = "aa"
    data1["b"] = "bb"
    data1["c"] = "cc"
    fmt.Println(data1)

 
 var data = map[string]map[string]string{}
    data["a"]["w"] = "x"
    data["b"]["w"] = "x"
    data["c"]["w"] = "x"
    fmt.Println(data) /* 

var data = map[string]map[string]string{
    "a": map[string]string{},
    "b": map[string]string{},
    "c": map[string]string{},
}

data["a"]["w"] = "x"
data["b"]["w"] = "x"
data["c"]["w"] = "x"
fmt.Println(data)*/


}
