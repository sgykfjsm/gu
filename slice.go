package gu

import (
	"reflect"
)

// Contains returns true if needle is found in haystack, otherwise it returns false.
// If needle is different data type from haystack's type, false will be returned always.
func Contains(needle, haystack interface{}) bool {
	h := reflect.ValueOf(haystack)
	if h.Kind() == reflect.Slice && h.Len() == 0 {
		return false
	} else if h.Kind() != reflect.Slice {
		return equals(needle, haystack)
	}

	islice := make([]interface{}, h.Len())
	for i := 0; i < h.Len(); i++ {
		islice[i] = h.Index(i).Interface()
	}

	needleType := reflect.TypeOf(needle)
	haystackType := reflect.TypeOf(islice[0])

	if needle != nil && needleType.Kind() != haystackType.Kind() {
		return false
	}

	n := reflect.ValueOf(needle).String()
	for i := range islice {
		if n == reflect.ValueOf(islice[i]).String() {
			return true
		}
	}
	return false
}

func equals(needle, haystack interface{}) bool {
	needleType := reflect.TypeOf(needle)
	haystackType := reflect.TypeOf(haystack)

	if needle != nil && needleType.Kind() != haystackType.Kind() {
		return false
	}

	if reflect.ValueOf(needle).String() == reflect.ValueOf(haystack).String() {
		return true
	}

	return false
}
