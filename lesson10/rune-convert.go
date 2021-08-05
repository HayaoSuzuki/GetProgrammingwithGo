package main

import (
	"fmt"
	"strconv"
)

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Println(string(pi), string(alpha), string(omega), string(bang))

	countdown := 10
	str := "発射まで" + strconv.Itoa(countdown) + "秒。"
	fmt.Println(str)

	countdown = 9
	str = fmt.Sprintf("発射まで%v秒。", countdown)
	fmt.Println(str)

	c, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("突然の死！")
	}
	fmt.Println(c)
}
