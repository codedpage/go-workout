package main

import "sort"
import "fmt"

type byLength []string

func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"ccc", "a", "dddd", "eeeee", "bb"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits) //[a bb ccc dddd eeeee]
}

