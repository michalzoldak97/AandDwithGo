package main

import (
	"fmt"

	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Creational/factory"
)

func testFactory() {
	s, err := factory.GetCar("Speedrunner100")
	if err != nil {
		fmt.Print(err)
		return
	}

	t, err := factory.GetCar("Truck3")
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Fast car is %v it has %v power and %v wheels\n", s.Model(), s.NumHP(), s.NumWheels())
	fmt.Printf("Truck car is %v it has %v power and %v wheels\n", t.Model(), t.NumHP(), t.NumWheels())
}

func main() {
	testFactory()
}
