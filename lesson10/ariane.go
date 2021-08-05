package main

import (
	"fmt"
	"math"
)

func main() {
	var bh float64 = 32768
	var h = int16(bh)
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		fmt.Println("爆発！")
	}
	fmt.Println(h)
}
