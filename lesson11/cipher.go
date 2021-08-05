package main

import (
	"fmt"
	"strings"
)

func charToDiff2(c byte) int {
	return int(c - 'A')
}

func shiftToRight(c byte, d int) byte {
	if c >= 'A' && c <= 'Z' {
		c += byte(d)
		if c > 'Z' {
			c -= 26
		}
	}
	return c
}

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"

	plainText = strings.Replace(plainText, " ", "", -1)
	plainText = strings.ToUpper(plainText)

	for i := 0; i < len(plainText); i++ {
		j := i % len(keyword)

		fmt.Printf("%c", shiftToRight(plainText[i], charToDiff2(keyword[j])))
	}
	fmt.Println("")
}
