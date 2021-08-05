package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func isLeap(year int) bool {
	var leap = (year%400 == 0) || ((year%4 == 0) && (year%100 != 0))
	return leap
}

func main() {
	for count := 0; count < 10; count++ {
		year := rand.Intn(9998) + 1
		month := rand.Intn(12) + 1
		daysInMonth := 31
		switch month {
		case 2:
			if isLeap(year) {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		day := rand.Intn(daysInMonth)
		fmt.Println(era, year, month, day)
	}
}
