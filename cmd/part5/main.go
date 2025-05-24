package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = "résumé"
	var indexed = myString[1]
	fmt.Printf("%v, %T \n", indexed, indexed)
	for i, v := range myString{
		fmt.Printf("index: %v and val: %v\n", i, v)
	}
	fmt.Printf("-----------------------\n")
	var myString2 = []rune("résumé")
	var indexer = myString2[1]
	fmt.Printf("%v, %T \n", indexer, indexer)
	for i, v := range myString2{
		fmt.Printf("index: %v and val: %v\n", i, v)
	}

	var myRune = 'a'
	fmt.Printf("myRun: %v \n", myRune)

	var strSlice = []string{"a", "d", "n", "o", "s"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("Concatenated String = %v\n", catStr)
}