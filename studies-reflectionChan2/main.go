package csvparse

import (
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
	"sync"
	"time"
)

// CSVDecoder holds the header field map and io.reader interface
type CSVDecoder struct {
	header map[string]int
	file   io.Reader
}

// NewDecoder allocates a new instance of CSVDecoder
func NewDecoder(input io.Reader) *CSVDecoder {
	return &CSVDecoder{
		header: make(map[string]int),
		file:   input,
	}
}

// DecodeCSV unmarshalls CSV file to a specified struct type
func (d *CSVDecoder) DecodeCSV(v interface{}) error {
	// Optional timer function for determining function duration
	// defer timeTrack(time.Now(), "DecodeCSVtoStruct")

	inChan, headerMap := ParseInout(d.file, v)

	// check interface type (v)
	// if type is not a slice, return error
	// slice := CheckInterfaceValue(v)
	slice := reflect.ValueOf(v).Elem()
	if slice.Kind() != reflect.Slice {
		return errors.New("Only slice types are permited")
	}
	// check inner slice type
	// if type is not a struct, return error
	innerTypePtr := reflect.TypeOf(v).Elem().Elem()
	innerTypeVal := reflect.TypeOf(v).Elem().Elem().Elem()
	if innerTypeVal.Kind() != reflect.Struct {
		return errors.New("Only slice of stucts type permited")
	}

	outChan := makeChannel(innerTypePtr, reflect.BothDir, 0)

	var wg sync.WaitGroup
	wg.Add(1000)

	go func() {
		wg.Wait()
		outChan.Close()
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			for csvRow := range inChan {
				innerValueRow := reflect.New(innerTypeVal)
				for j := 0; j < innerValueRow.Elem().NumField(); j++ {
					name := reflect.Indirect(innerValueRow).Type().Field(j).Name
					switch innerValueRow.Elem().Type().Field(j).Type {
					case reflect.TypeOf(time.Now()):
						if _, ok := headerMap[name]; ok {
							val := ParseDate(csvRow[headerMap[name]])
							innerValueRow.Elem().FieldByName(name).Set(reflect.ValueOf(val))
						}
					default:
						if format, ok := reflect.Indirect(innerValueRow).Type().Field(j).Tag.Lookup("fmt"); ok {
							if format != "-" {
								fmtvalue, err := FormatStringVals(format, fmt.Sprint(innerValueRow.Elem().FieldByName(name)))
								if err != nil {
									log.Fatalln(err)
								}
								innerValueRow.Elem().FieldByName(name).Set(reflect.ValueOf(fmtvalue))
							}
						}
					}
				}
				outChan.Send(innerValueRow)
			}
		}()
	}

	for ok := true; ok; {
		var element reflect.Value
		if element, ok = outChan.Recv(); ok {
			slice.Set(reflect.Append(slice, element))
		}
	}
	return nil
}
