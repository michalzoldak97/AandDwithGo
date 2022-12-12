package main

import (
	"fmt"

	adapter "github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/Adapter"
	plane "github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/Bridge"
	composite "github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/Composite"
)

func runCompositeExamle() {
	s := &composite.School{}
	s.CreateClasses()
	nameToSearch := "Lin"
	hasName, name := s.SearchName(nameToSearch)
	fmt.Printf("There is name %v\n full is %v\n", hasName, name)
}

func runAdapterExample() {
	u := &adapter.User{}
	dellDisplay := &adapter.DPDisplay{}
	acerDisplay := &adapter.DVIDisplay{}

	dviDpAdapter := &adapter.DVIAdapter{
		DviDisplay: acerDisplay,
	}

	u.PlugInDisplayPortIntoDisplay(dellDisplay)
	u.PlugInDisplayPortIntoDisplay(dviDpAdapter)
}

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

	runAdapterExample()
	runBridgeExample()
	runCompositeExamle()
}
