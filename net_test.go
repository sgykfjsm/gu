package gu

import "testing"

func TestFindLocalIPv4Address(t *testing.T) {
	result, err := FindLocalIPv4Address()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == "" {
		t.Error("empty return is not expected")
	}
}
