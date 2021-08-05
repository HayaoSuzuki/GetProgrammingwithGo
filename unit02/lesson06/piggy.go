package main

import (
	"fmt"
	"math/rand"
)

func getCoin() float64 {
	var coin = 0.0
	if n := rand.Intn(3); n == 0 {
		coin = 0.05
	} else if n == 1 {
		coin = 0.10
	} else {
		coin = 0.25
	}
	return coin
}

func main() {
	const target = 20.0
	var piggy = 0.0
	for piggy <= target {
		piggy += getCoin()
		fmt.Printf("%5.2f\n", piggy)
	}
}
