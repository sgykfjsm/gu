package gu

import (
	"reflect"
)

// Contains returns true if needle is found in haystack, otherwise it returns false.
func Contains(needle, haystack interface{}) bool {
	h := reflect.ValueOf(haystack)
	if h.Kind() != reflect.Slice || h.Len() == 0 {
		return false
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

	nv := reflect.ValueOf(needle).String()
	for i := range islice {
		if nv == reflect.ValueOf(islice[i]).String() {
			return true
		}
	}
	return false
}
