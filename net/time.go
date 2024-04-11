package net

import (
	"fmt"
	"time"
)

func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("use time = %v\n", tc)
	}
}
