package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 実行のたびに毎回変化してほしい
	var answer = rand.Intn(100) + 1
	for {
		var guess = rand.Intn(100) + 1
		if guess == answer {
			fmt.Printf("正解は %v でした！\n", guess)
			break
		} else if guess > answer {
			fmt.Printf("%v は正解よりも大きいようです。\n", guess)
		} else {
			fmt.Printf("%v は正解よりも小さいようです。\n", guess)
		}
	}
}
