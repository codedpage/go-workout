package main

import "fmt"

func main() {

	b := [4]interface{}{1, "hello", 2.5, nil}
	infPrint(b)
}

func infPrint(b [4]interface{}) {
	for k, v := range b {

		switch v.(type) {
		case string:
			fmt.Println(k, "=====", v.(string))

		case int:
			fmt.Println(k, "=====", v.(int))

		case float64:
			fmt.Println(k, "=====", v.(float64))

		case bool:
			fmt.Println(k, "=====", v.(bool))

		case nil:
			fmt.Println(k, "=====", nil)

		}
	}
}

//https://go.dev/play/p/xpdp3OQWVex