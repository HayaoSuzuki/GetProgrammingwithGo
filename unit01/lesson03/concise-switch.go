package main

import "fmt"

func main() {
	fmt.Println("洞窟の入口だ。東に進む道もある。")
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("君は、更に山に登る。")
	case "enter cave", "go inside":
		fmt.Println("君は薄暗い洞窟の中にいる。")
	case "read sign":
		fmt.Println("「未成年立ち入り禁止」と書いてある。")
	default:
		fmt.Println("なんだか、よくわからない。")
	}
}
