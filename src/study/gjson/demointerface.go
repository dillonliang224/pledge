package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	fmt.Println(f, err, reflect.TypeOf(f))

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, " is string", vv)
		case float64:
			fmt.Println(k, " is float64", vv)
		case []interface{}:
			fmt.Println(k, " is array: ")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, " unknown")
		}

	}
}
