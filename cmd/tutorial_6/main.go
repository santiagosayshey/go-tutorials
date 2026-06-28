package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// (e gasEngine) assigns it to the gasEngine struct and e allows the function to access the fields from it
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

// the point of an interface is to define what a function needs
// without tying it to a specific struct.
// any struct that has these methods can be used,
// regardless of what else it has.
// the function signature MUST match exactly. parameters and return value
type engine interface {
	milesLeft() uint8
}

func canMakeit(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myGasEngine gasEngine = gasEngine{25, 15}
	var myElectricEngine electricEngine = electricEngine{10, 20}
	canMakeit(myGasEngine, 20)
	canMakeit(myElectricEngine, 10)
}
