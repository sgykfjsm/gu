package gu

import (
	"testing"
	"reflect"
	"fmt"
)

func TestIsContain(t *testing.T) {
	testCases := []struct {
		Needle   interface{}
		HayStack interface{}
		Expect bool
	}{
		{true, []bool{true, false}, true},
		{true, true, true},
		{true, []bool{}, false},
		{true, []string{"true", "false"}, false},
		{"string", "string", true},
		{"string", []string{"this", "is", "string", "slice"}, true},
		{"string", []string{"this", "is", "slice"}, false},
		{"string", []bool{true, false}, false},
		{5, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{5, 5, true},
		{1.0, []float64{0.1, 1.0, 2.2}, true},
		{nil, nil, true},
		{nil, []bool{true, false}, false},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%s/%t", reflect.ValueOf(tc.Needle).Kind().String(), tc.Expect)
		t.Run(testName, func(t *testing.T) {
			result := Contains(tc.Needle, tc.HayStack)
			if result != tc.Expect {
				t.Error(fmt.Sprintf("%+v in %q: expected %t but not.", tc.Needle, tc.HayStack, tc.Expect))
			}
		})
	}
}
