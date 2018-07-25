package main

import (
	"fmt"
	"reflect"
	"time"
)

type record struct {
	Name  string    `csv:"FullName" json:"json-FullName"`
	Date  time.Time `csv:"Time" json:"Time"`
	Count int       `csv:"ZipCode" json:"json-ZipCode"`
	Paid  bool      `csv:"Loan Amout" json:"json-Amount"`
}

func main() {
	r := record{}
	fmt.Println("Before", r)
	objInfo(&r)
	fmt.Println("After", r)
}

func objInfo(v interface{}) {
	for i := 0; i < reflect.ValueOf(v).Elem().NumField(); i++ {
		fldName := reflect.Indirect(reflect.ValueOf(v)).Type().Field(i).Name
		fldValue := reflect.ValueOf(v).Elem().Field(i)
		fldType := reflect.ValueOf(v).Elem().Field(i).Type().Name()
		fldTag := reflect.Indirect(reflect.ValueOf(v)).Type().Field(i).Tag
		fldTagcsvVal := reflect.Indirect(reflect.ValueOf(v)).Type().Field(i).Tag.Get("csv")
		fldTagjsonVal := reflect.Indirect(reflect.ValueOf(v)).Type().Field(i).Tag.Get("json")
		fldIndex := reflect.Indirect(reflect.ValueOf(v)).Type().Field(i).Index[0]
		switch fldName {
		case "Name":
			fmt.Println(reflect.ValueOf(v).Elem().FieldByName(fldName))
			reflect.ValueOf(v).Elem().FieldByName(fldName).Set(reflect.ValueOf("HelloWorld"))
		case "Count":
			reflect.ValueOf(v).Elem().FieldByName(fldName)
			reflect.ValueOf(v).Elem().FieldByName(fldName).Set(reflect.ValueOf(2))
		case "Paid":
			reflect.ValueOf(v).Elem().FieldByName(fldName)
			reflect.ValueOf(v).Elem().FieldByName(fldName).Set(reflect.ValueOf(true))
		}

		fmt.Println(fldIndex, fldName, fldValue, fldType, fldTag, fldTagcsvVal, fldTagjsonVal)
	}
	// fmt.Println("Type", reflect.Indirect(reflect.ValueOf(v)).Type())
	fmt.Println("ValueA", reflect.ValueOf(v).Elem().Type())
	fmt.Println("ValueB", reflect.ValueOf(v).Elem())
}
