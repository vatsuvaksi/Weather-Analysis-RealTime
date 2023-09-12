package structutils

import (
	"fmt"
	"reflect"
)

func PrintStructFields(s interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		fmt.Println("Not a struct")
		return
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fieldValue := field.Interface()
		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}
}
