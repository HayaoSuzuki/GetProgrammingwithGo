package main

import "fmt"

func main() {
	// マラカンドラは、それよりずっと近距離にある。約二十八日間で到着するだろう。
	// マラカンドラはC.S ルイス『別世界物語』に登場する火星の別名
	const distance = 56000000 // km
	const days = 28
	const hoursPerDay = 24
	fmt.Println(distance/(days*hoursPerDay), "km/h")
}
