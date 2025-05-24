package main

import "fmt"

type gasEngine struct{
	mpg uint16
	gallons uint16
}

type electricEngine struct{
	mpkwh uint16
	kwh uint16
}

func (e electricEngine) milesLeft() uint16{
	return e.kwh * e.mpkwh
}

type engine interface{
	milesLeft() uint16
}

func canMakeIt(e engine, miles uint16) {
	if miles <= e.milesLeft(){
		fmt.Printf("You can make it there\n")
	}else{
		fmt.Printf("Need to fuel it up first\n")
	}
}

func (e gasEngine) milesLeft() uint16{
	return e.gallons * e.mpg
}

func main(){
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 40}
	fmt.Printf("Mile my engine can cover : %v \n", myEngine.milesLeft())
	fmt.Printf("Can I make it to 1200 miles? \n")
	canMakeIt(myEngine, 1200)
	var myEngine2 electricEngine = electricEngine{kwh: 200, mpkwh: 10}
	fmt.Printf("Miles my engine can cover : %v \n", myEngine2.milesLeft())
	fmt.Printf("Can I make it to 1200 miles? \n")
	canMakeIt(myEngine2, 1200)
}