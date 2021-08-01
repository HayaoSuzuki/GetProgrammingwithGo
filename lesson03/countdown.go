package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Float64() <= 0.1 {
			fmt.Println("FAIL!")
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("Liftoff!")
	}
}
