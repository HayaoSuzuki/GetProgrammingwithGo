package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792) // km/s
	secondPerDay := big.NewInt(86400)
	dayPerYear := big.NewInt(365)
	distance := new(big.Int)
	distance.SetString("236000000000000000", 10) // km

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondPerDay)

	lightYears := new(big.Int)
	lightYears.Div(days, dayPerYear)

	fmt.Println("おおいぬ座矮小銀河まで、", lightYears, "光年かかる。")
}
