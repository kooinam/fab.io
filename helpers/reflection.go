package helpers

import (
	"fmt"
	"reflect"
)

func GetFieldNamesByKey(s interface{}, key string) []string {
	fieldNames := []string{}
	rt := reflect.TypeOf(s).Elem()

	if rt.Kind() != reflect.Struct {
		panic(fmt.Sprintf("bad type - %v", rt.Kind()))
	}

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag.Get(key)

		fmt.Printf("%v - %v\n", field.Name, tag)

		if len(tag) > 0 {
			fieldNames = append(fieldNames, field.Name)
		}
	}

	return fieldNames
}
