package main

import "fmt"

func main() {
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("君は薄暗い洞窟の中にいる。")
	case room == "lake":
		fmt.Println("硬そうに氷が張っている。")
		fallthrough
	case room == "underwater":
		fmt.Println("水は凍るくらいに冷たい。")
	}
}
