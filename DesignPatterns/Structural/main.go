package main

import (
	"fmt"

	plane "github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/Bridge"
)

func runBridgeExample() {
	aesaRadar := &plane.AESA{}
	parRadar := &plane.PAR{}

	f16 := &plane.Interceptor{}
	f16.SetRadar(aesaRadar)

	b1 := &plane.Bomber{}
	b1.SetRadar(parRadar)

	f16.StartRadar()
	b1.StartRadar()
}

func main() {
	fmt.Println("Structural design patterns")
	runBridgeExample()
}
