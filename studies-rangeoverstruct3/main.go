package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Product struct {
		Name  string
		Price string
	}

	x := Product{}
	newStruct(x)
}

func newStruct(v interface{}, w Product) {
	vType := reflect.TypeOf(v)       // this type of this variable is reflect.Type
	vPointer := reflect.New(vType)   // this type of this variable is reflect.Value.
	vValue := vPointer.Elem()        // this type of this variable is reflect.Value.
	vInterface := vValue.Interface() // this type of this variable is interface{}
	v2 := vInterface.(struct)       // this type of this variable is product

	v2.Name = "Toothbrush"
	v2.Price = "2.50"

	fmt.Println(v2.Name)
	fmt.Println(v2.Price)
}
