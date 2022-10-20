package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
)

func main() {

resp, err := http.Get("http://localhost/php_test/1.php?a=ram") //1.php => echo ($_GET['a']);
    if err != nil {
         fmt.Println(err)
    }
 
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Println(contents) //ascii values =>[114 97 109 32]
	fmt.Printf("%s\n", string(contents)) //ascii values to string =>ram

}