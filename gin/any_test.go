package main

import (
	"fmt"
	"testing"
	"time"
)

func TestLocalTime(t *testing.T) {
	tz, err := time.LoadLocation("America/New_York")
	fmt.Println("tz", tz, err)

	fmt.Println("time.Now().In(America/New_York)", time.Date(2024, 3, 10, 3, 1, 0, 0, tz).String())
	fmt.Println("time.Now().In(America/New_York) is DST", time.Date(2024, 3, 10, 3, 0, 0, 0, tz).IsDST())
	fmt.Println("time.Now().In(America/New_York)", time.Date(2024, 3, 10, 2, 59, 0, 0, tz).String())
	fmt.Println("time.Now().In(America/New_York) is DST", time.Date(2024, 3, 10, 2, 59, 0, 0, tz).IsDST())
}
