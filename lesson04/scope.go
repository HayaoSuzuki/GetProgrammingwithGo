package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	for count < 10 { // 新しいスコープ開始
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	} // スコープ修了
	// fmt.Println(num)  // コンパイルエラー
}
