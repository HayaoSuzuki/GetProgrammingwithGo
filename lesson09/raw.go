package main

import "fmt"

func main() {
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with the \n escape sequence`)
	fmt.Println(`
	peace be upon you
	upon you be peace`)
	fmt.Printf("%v は%[1]T 型です\n", "文字列リテラル")
	fmt.Printf("%v は%[1]T 型です\n", `raw文字列リテラル`)
}
