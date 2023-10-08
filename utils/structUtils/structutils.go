package structutils

import (
	"fmt"
	"real-time-weather-app/utils/loggers"
	"reflect"

	"github.com/sirupsen/logrus"
)

func PrintStructFields(s interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		loggers.Logger.WithField("Key", logrus.Fields{
			"Input":   s,
			"isFatal": false,
		}).Warn("Input is not a struct")
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
