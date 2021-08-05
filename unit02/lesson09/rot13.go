package main

import "fmt"

func main() {
	var message = "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ {
		var c = message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
}
