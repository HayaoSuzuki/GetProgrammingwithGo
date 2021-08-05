package main

import "fmt"

func charToDiff(c byte) int {
	return int(c - 'A')
}

func shiftToLeft(c byte, d int) byte {
	if c >= 'A' && c <= 'Z' {
		c -= byte(d)
		if c < 'A' {
			c += 26
		}
	}
	return c
}

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	//cipherText := "ECFRZKYGLGRMUSDHRXK"
	keyword := "GOLANG"

	for i := 0; i < len(cipherText); i++ {
		j := i % len(keyword)

		fmt.Printf("%c", shiftToLeft(cipherText[i], charToDiff(keyword[j])))
	}
	fmt.Println("")
}
