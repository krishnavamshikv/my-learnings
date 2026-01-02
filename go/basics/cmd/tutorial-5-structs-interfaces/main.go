package main

import "fmt"

// structs and interfaces

// we can use anothe struct inside a struct as well and use its method also
type gasEngine struct {
	mpg     uint8
	gallons uint8
	start   // another struct
}

type start struct {
	say string
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// methods for stuct
func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// both sturcts now use same method (milesLeft) in this case we can use interface
type engine interface {
	milesLeft() uint8 // any struct that has milesLeft() method
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("need to fuel first!")
	}
}

func main() {
	myEngine := gasEngine{25, 20, start{"Car started bro"}}
	electicEngine := electricEngine{25, 20}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.say)
	canMakeIt(myEngine, 200)
	canMakeIt(electicEngine, 200)
}
