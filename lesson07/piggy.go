package main

import (
	"fmt"
	"math/rand"
)

func getCoin() int {
	var coin = 0
	if n := rand.Intn(3); n == 0 {
		coin = 5
	} else if n == 1 {
		coin = 10
	} else {
		coin = 25
	}
	return coin
}

func main() {
	const target = 2000 // cent
	var piggy = 0       // cent
	for piggy <= target {
		piggy += getCoin()
		q, r := piggy/100, piggy%100
		fmt.Printf("$%v.%v\n", q, r)
	}
}
