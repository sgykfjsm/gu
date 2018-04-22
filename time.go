package gu

import "time"

func ElapseSecond(start time.Time) float64 {
	return float64(time.Since(start)) / float64(time.Second)
}
