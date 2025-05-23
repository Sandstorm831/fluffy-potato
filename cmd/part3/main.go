package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "hello world"
	printMe(printValue)

	var num int = 10
	var den int = 0
	var result, remainder, err = intDivisor(num, den)
	if err != nil {
		fmt.Printf("We got some error : %v", err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with reaminder of %v", result, remainder)
	}
}

func printMe(pval string) {
	fmt.Printf("the printerMe is called >>> %v", pval)
}

func intDivisor(num int, den int) (int, int, error) {
	var err error
	if den == 0 {
		err = errors.New("Can't divide by zero")
		return 0, 0, err
	} else {
		var q int = num / den
		var r int = num % den
		return q, r, err
	}
}
