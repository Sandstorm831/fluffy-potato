package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 2341
	intNum = intNum + 1
	fmt.Println(intNum)
	var a1 int = 3
	var a2 int = 2
	fmt.Println(a1 / a2)
	fmt.Println(a1 % a2)
	var gamma string = "γ"
	fmt.Println(len(gamma))                    // Print number of bytes
	fmt.Println(utf8.RuneCountInString(gamma)) // Print the length of string

	var myrune rune = 'a'
	fmt.Println(myrune)

	const pi float32 = 3.141592
	fmt.Println(pi)
}
