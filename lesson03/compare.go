package main

import "fmt"

func main() {
	fmt.Println("入り口の近くに「未成年立ち入り禁止」という立札がある。")
	var age = 41
	var minor = age < 18
	fmt.Printf("%v 歳のぼくは、未成年か？ %v", age, minor)
}
