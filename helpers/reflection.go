package helpers

import (
	"fmt"
	"reflect"
	"strings"
)

func GetFieldsDetailsByTag(s interface{}, tag string) []*Dictionary {
	fieldsDetails := []*Dictionary{}
	rt := reflect.TypeOf(s).Elem()

	if rt.Kind() != reflect.Struct {
		panic(fmt.Sprintf("bad type - %v", rt.Kind()))
	}

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tagValue := strings.Split(field.Tag.Get("fab"), ",")[0]

		if len(tagValue) > 0 {
			fieldDetails := MakeDictionary(H{
				"name":  field.Name,
				"value": strings.Split(tagValue, ":")[1],
			})
			fieldsDetails = append(fieldsDetails, fieldDetails)
		}
	}

	return fieldsDetails
}

func GetFieldValueByName(s interface{}, name string) interface{} {
	rt := reflect.ValueOf(s).Elem()

	if rt.Kind() != reflect.Struct {
		panic(fmt.Sprintf("bad type - %v", rt.Kind()))
	}

	field := rt.FieldByName(name)

	return field.Interface()
}

func SetFieldValueByNameStr(s interface{}, name string, value string) {
	rt := reflect.ValueOf(s).Elem()
	field := rt.FieldByName(name)

	field.SetString(value)
}
