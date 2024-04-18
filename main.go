package main

import "fmt"

func main() {
	currencyMap := GetCurrencies()

	fmt.Println(currencyMap)
	//this func going to get currencies data everyday at 15.32
	go func() {
		ScheduleByTime(func() {
			currencyMap = GetCurrencies()
		}, 15, 32)
	}()

}
