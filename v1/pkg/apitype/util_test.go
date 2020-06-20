package apitype_test

import "reflect"

func fieldCount(val interface{}) int {
	return reflectFieldCount(reflect.TypeOf(val))
}

func reflectFieldCount(t reflect.Type) int {
	count := 0

	for j := 0; j < t.NumField(); j++ {
		if t.Field(j).Anonymous {
			count += reflectFieldCount(t.Field(j).Type)
		} else {
			count++
		}
	}

	return count
}
