// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import  ("io/ioutil"
  "math/rand"
  "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    //ele := []byte("")
	var countryCapitalMap map[string]string
	/* create a map*/
	countryCapitalMap = make(map[string]string)

	/* insert key-value pairs in the map*/
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	c := 1
	for country := range countryCapitalMap {
		ele := []byte(country)
		fn := fmt.Sprintf("file_%d%d.txt", c,rand.Intn(50))
		print(fn)
		ioutil.WriteFile(fn, ele, 0644)
		c++
	}

   

 
 



}
