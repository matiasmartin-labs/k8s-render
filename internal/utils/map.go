package utils

import (
	"reflect"

)

func StructToMap(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := typ.Field(i)
		value := v.Field(i)

		if field.Type.Kind() == reflect.Func {
			// Skip functions
			continue
		}

		if field.Type.Kind() == reflect.Struct {
			result[field.Name] = StructToMap(value.Interface())
		} else {
			result[field.Name] = value.Interface()
		}

		if field.Type.Kind() == reflect.Slice {
			slice := make([]interface{}, value.Len())
			for j := 0; j < value.Len(); j++ {
				slice[j] = value.Index(j).Interface()
			}
			result[field.Name] = slice
		}

		if field.Type.Kind() == reflect.Map {
			mapValue := make(map[string]interface{})
			for _, key := range value.MapKeys() {
				mapValue[key.String()] = value.MapIndex(key).Interface()
			}
			result[field.Name] = mapValue
		}

		if field.Type.Kind() == reflect.Array {
			array := make([]interface{}, value.Len())
			for j := 0; j < value.Len(); j++ {
				array[j] = value.Index(j).Interface()
			}
			result[field.Name] = array
		}

		if field.Type.Kind() == reflect.Interface {
			if value.IsNil() {
				result[field.Name] = nil
			} else {
				result[field.Name] = value.Interface()
			}
		}

	}
	return result
}
