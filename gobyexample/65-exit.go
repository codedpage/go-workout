package main

import "fmt"
import "os"

func main() {

    defer fmt.Println("!")

    os.Exit(3)
}

/*
$ go run 65-exit.go  echo $?
*/