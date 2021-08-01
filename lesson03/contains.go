package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("君は薄暗い洞窟の中にいる。")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	var wearShades = strings.Contains(command, "glasses")
	var readText = strings.Contains(command, "read")
	fmt.Println("洞窟を出る：", exit)
	fmt.Println("サングラスをかける：", wearShades)
	fmt.Println("立札を呼んだ：", readText)
}
