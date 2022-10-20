package main
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
	"time"
)

func channel_hello(done chan bool) { 

resp, err := http.Get("http://localhost/php_test/1.php?a=ram")
    if err != nil {
         fmt.Println(err)
    }
 
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("%T", resp.Body) //*http.bodyEOFSignal
	fmt.Println(contents) //[114 97 109 32]
	fmt.Printf("%s\n", string(contents)) //ram
	done <- true //return true
}

 
func main() {  
    done := make(chan bool)
    go channel_hello(done) 	
    time.Sleep(1 * time.Second)  //Or <-done  
            
    fmt.Println("main")
}