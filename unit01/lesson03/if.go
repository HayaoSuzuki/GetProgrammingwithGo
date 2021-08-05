package main

import "fmt"

func main() {
	var command = "go east"

	if command == "go east" {
		fmt.Println("君は、更に山に登る。")
	} else if command == "go inside" {
		fmt.Println("君は洞窟に入り、そこで一生を過ごす。")
	} else {
		fmt.Println("なんだか、よくわからない。")
	}
}
