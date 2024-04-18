package main

import (
	"fmt"
	"math"
	"time"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func ScheduleByTime(call func(), hour, minute int) {
	for {
		now := time.Now()

		targetTime := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())

		now = time.Now()

		if now.After(targetTime) {
			targetTime = targetTime.Add(24 * time.Hour)
		}

		waitTime := targetTime.Sub(now)

		fmt.Printf("current time is: %v and %v hours remaining to e-archiving...\n", now, waitTime.Hours())

		time.Sleep(waitTime)

		now = time.Now()
		fmt.Println("call running!", now)
		call()
	}
}
