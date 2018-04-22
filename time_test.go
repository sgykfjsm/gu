package gu

import (
	"testing"
	"time"
)

func TestElapseSecond(t *testing.T) {
	start := time.Now()
	time.Sleep(time.Second)
	expected := 1
	actual := int(ElapseSecond(start))
	if actual != expected {
		t.Errorf("expected %v but actual %v", expected, actual)
	}
}
