package main

import (
	"fmt"
	"reflect"
)

func SetKeyValue(keyupdate string, mp map[string]interface{}, newvalue interface{}) {
	// If a key exists in a map, update its value
	if _, exists := mp[keyupdate]; exists == true {
		mp[keyupdate] = newvalue
	}

	for key, val := range mp {
		v := reflect.ValueOf(val)
		fmt.Println(key, val, "  kind=", v.Kind())

		switch v.Kind() {

		// If it's a slice, iterate over its elements
		case reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				element := v.Index(i).Interface()
				// fmt.Println("wdfsdf",reflect.TypeOf(element))
				if subMap, ok := element.(map[string]interface{}); ok {
					SetKeyValue(keyupdate, subMap, newvalue)
				}
			}

		//If it's a map , recursively run same function again for that map
		case reflect.Map:
			SetKeyValue(keyupdate, val.(map[string]interface{}), newvalue)

		}

	}
}

func main() {
	var m = map[string]interface{}{
		"Name": "Vaibhav Bhardwaj",
		"DOB":  24 - 01 - 2002,
		"city": "Delhi",
		"pin":  110075,
		// field named "Address" and assigns it a slice ([]interface{}). The square brackets [] indicate that it's a slice, and interface{} allows elements of any data type to be stored in the slice.
		"NewTest": map[string]interface{}{
			"street":  "Testtt",
			"plot_no": 96,
			"city":    "Example city",
			"pin":     633078,
		},
		"Address": []interface{}{
			//nside the slice, there are two elements, each represented by a map.

			map[string]interface{}{
				"street":  "Ashirvad chowk",
				"plot_no": 26,
				"city":    "Dwarka",
				"pin":     110078,
			},
			map[string]interface{}{
				"street":  "Lovely Chowk",
				"plot_no": 26,
				"city":    "London",
				"pin":     923478,
			},
		},
		"Salary":      800000,
		"Designation": "Developer",
	}
	//
	var val interface{}
	val = "New York"

	SetKeyValue("city", m, val)
	fmt.Println(m)

}
