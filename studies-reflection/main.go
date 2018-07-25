package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type A struct {
	Name string `column:"email"`
}

func main() {
	bbb(&A{})
}

func aaa(v interface{}) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Slice {
		t = t.Elem()
	} else {
		panic("Input param is not a slice")
	}
	slice := reflect.ValueOf(v)
	if t.Kind() == reflect.Ptr {
		slice = slice.Elem()
	}
	sliceofStruct := slice.Type()
	fmt.Printf("Outer Type %s:\n", sliceofStruct)

	sliceElementType := sliceofStruct.Elem()
	if sliceElementType.Kind() == reflect.Ptr {
		sliceElementType = sliceElementType.Elem()
	}
	fmt.Printf("Inner Type %v:\n", sliceElementType)

	for i := 0; i < 5; i++ {
		newitem := reflect.New(sliceElementType)
		newitem.Elem().FieldByName("Name").SetString(fmt.Sprintf("Grzes %d", i))

		s := newitem.Elem()
		for i := 0; i < s.NumField(); i++ {
			col := s.Type().Field(i).Tag.Get("column")
			fmt.Println(col, s.Field(i).Addr().Interface())
		}

		slice.Set(reflect.Append(slice, newitem))
	}
}

func bbb(v interface{}) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("Input param is not a struct")
	}

	models := reflect.New(reflect.SliceOf(reflect.TypeOf(v))).Interface()

	aaa(models)

	fmt.Println(models)

	b, err := json.Marshal(models)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
