package helpers

import (
	"fmt"
	"strconv"
)

// Dictionary used to provide accessibilities to dynamic hash
type Dictionary struct {
	properties H
}

// MakeDictionary used to instantiate dictionary instance
func MakeDictionary(properties H) *Dictionary {
	return &Dictionary{
		properties: properties,
	}
}

// Value used to retrieve params value
func (dict *Dictionary) Value(key string) interface{} {
	value := dict.properties[key]

	return value
}

// ValueStr used to retrieve params value in string
func (dict *Dictionary) ValueStr(key string) string {
	value := dict.properties[key]

	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}

// ValueInt used to retrieve params value in int
func (dict *Dictionary) ValueInt(key string, fallback int) int {
	value := dict.properties[key]

	if value == nil {
		return fallback
	}

	switch value.(type) {
	case string:
		i, err := strconv.Atoi(value.(string))

		if err != nil {
			return fallback
		}

		return i
	case float64:
		return int(value.(float64))
	}

	return value.(int)
}

// ValueFloat64 used to retrieve params value in float64
func (dict *Dictionary) ValueFloat64(key string, fallback float64) float64 {
	value := dict.properties[key]

	if value == nil {
		return fallback
	}

	switch value.(type) {
	case string:
		f, err := strconv.ParseFloat(value.(string), 64)

		if err != nil {
			return fallback
		}

		return f
	}

	return value.(float64)
}

// Set used to set property
func (dict *Dictionary) Set(key string, value interface{}) {
	dict.properties[key] = value
}
