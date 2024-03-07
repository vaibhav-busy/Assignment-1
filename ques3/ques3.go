package main

import (
	"fmt"
	"reflect"
)

// PopulateStruct function populates fields of a struct using data from a map.
func PopulateStruct(data map[string]interface{}, persn interface{}) {
	// Get the reflection value of the persn interface and navigate to its underlying struct
	resultValue := reflect.ValueOf(persn).Elem()

	// Iterate over each key-value pair in the data map
	for key, value := range data {
		// Get the field in the struct corresponding to the key
		persnfield := resultValue.FieldByName(key)

		// Check if the field is valid (exists in the struct) . Only do somethNing if person struct also has the same field
		if persnfield.IsValid() {
			// If the field is a struct itself, recursively populate its fields
			if persnfield.Kind() == reflect.Struct {
				// Check if the value is a map, indicating nested struct data
				if nestedMap, ok := value.(map[string]interface{}); ok {
					// Create a new instance of the nested struct type
					nestedStruct := reflect.New(persnfield.Type()).Interface()
					// Recursively populate the nested struct fields
					PopulateStruct(nestedMap, nestedStruct)
					// Set the populated nested struct back to the field
					persnfield.Set(reflect.ValueOf(nestedStruct).Elem())
				}
			} else {
				// For non-struct fields, set the value directly
				persnfield.Set(reflect.ValueOf(value))
			}
		}
	}

}

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	City  string
	State string
}

func main() {
	// Sample data to populate the Person struct
	data := map[string]interface{}{
		"Name":    "Vaibhav",
		"Age":     22,
		"pincode": 250002, // This field won't be used as it doesn't exist in the struct
		"Address": map[string]interface{}{
			"City":  "Meerut",
			"State": "Uttar Pradesh",
		},
	}

	// Create a pointer to a Person struct
	var personPtr *Person = &Person{}
	// Populate the Person struct fields using the data
	PopulateStruct(data, personPtr)

	// Print the populated Person struct
	fmt.Printf("%+v\n", *personPtr)
}
