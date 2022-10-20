package main

import "fmt"

func main() {

    var data1 = map[string]string{}
    data1["a"] = "aa"
    data1["b"] = "bb"
    data1["c"] = "cc"
    fmt.Println(data1)
 
    var data11 = map[string][]string{}
    data11["a"] = append(data11["a"], "aaa")
    data11["b"] = append(data11["b"], "bbb")
    data11["c"] = append(data11["c"], "ccc")
    fmt.Println(data11)

 
    //var s = map[int]string{}
    s := map[int]string{
            1:"a",
            2:"b",
            3:"c",
            }
    s[4] = "d"
    var data12 = map[string]map[int]string{}
    data12["a"] = s
    data12["b"] = s
    data12["c"] = s
    fmt.Println(data12)
 

var data2 = map[string]map[string]string{
    "a": map[string]string{
        "w": "x"},
    "b": map[string]string{
        "w": "x"},
    "c": map[string]string{
        "w": "x"},
    "d": map[string]string{},
}
fmt.Println(data2)

 
/*
var data3 = map[string]map[string]string{
    "a": map[string]string{},
    "b": map[string]string{},
    "c": map[string]string{},
}

data1["a"]["w"] = "x"
data1["b"]["w"] = "x"
data1["c"]["w"] = "x"
fmt.Println(data1) */



}
