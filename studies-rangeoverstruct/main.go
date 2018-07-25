package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type MyStruct struct {
	A, B, C string
	I       int
	D       string
	J       time.Time
}

func main() {
	ms := &MyStruct{"Green ", " Eggs", " and ", 2, " Ham      ", time.Now()}
	// Print it out now so we can see the difference
	fmt.Printf("%s%s%s%d%s%v\n", ms.A, ms.B, ms.C, ms.I, ms.D, ms.J)

	// We need a pointer so that we can set the value via reflection
	msValuePtr := reflect.ValueOf(ms)
	msValue := msValuePtr.Elem()

	for i := 0; i < msValue.NumField(); i++ {
		field := msValue.Field(i)
		// Ignore fields that don't have the same type as a string
		if field.Type() != reflect.TypeOf("") {
			continue
		}
		str := field.Interface().(string)
		str = strings.TrimSpace(str)
		field.SetString(str)
	}

	// We need a pointer so that we can set the value via reflection
	fmt.Printf("%s%s%s%d%s%v\n", ms.A, ms.B, ms.C, ms.I, ms.D, ms.J)
}
