package util

import (
	"math/rand"
	"time"
)

func GetRandomDuration(minValue, maxValue time.Duration) time.Duration {
	if minValue >= maxValue {
		return minValue
	}
	offset := time.Duration(rand.Int63n(int64(maxValue-minValue) + 1))
	return minValue + offset
}
