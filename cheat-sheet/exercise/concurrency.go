package main
import (
    "fmt"
)
func producer(ch chan int) {  
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Println("wrote-----", i, "to ch")
    }
    close(ch)
}
func main() {  
    ch := make(chan int, 5)
    go producer(ch)
    for v := range ch {
        fmt.Println("------read ", v,"from ch")
    }

	/*
		for {
			v, ok := <-ch
			if ok == false {
				break
			}
			fmt.Println("Received ", v, ok)
		}

	*/
}
//////////////////////
