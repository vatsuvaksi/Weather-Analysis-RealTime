package structutils

import (
	"fmt"
	"reflect"
)

func PrintStructFields(i interface{}) {
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		t := reflect.TypeOf(i)
		v := reflect.ValueOf(i)
		fmt.Println("\n", t.Name())
		for j := 0; j < v.NumField(); j++ {
			switch v.Field(j).Kind() {
			case reflect.Struct:
				printWeatherDetails(v.Field(j).Interface())
			case reflect.String:
				fmt.Printf("%s : %s\n", t.Field(j).Name, v.Field(j).String())
			case reflect.Float64:
				fmt.Printf("%s : %f\n", t.Field(j).Name, v.Field(j).Float())
			default:
				fmt.Println("Type not supported..")
			}
		}
	}
}

