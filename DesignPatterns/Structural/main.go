package main

import (
	"fmt"

	adapter "github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/Adapter"
	plane "github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/Bridge"
	composite "github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/Composite"
	decorator "github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/Decorator"
)

func runDecoratorExample() {
	suv := &decorator.Car{}
	suv.SetPrice(198000)

	suvWithTurbo := &decorator.Turbo{}
	suvWithTurbo.SetBaseOffer(suv)

	suvWithTurboAndAwd := &decorator.Awd{}
	suvWithTurboAndAwd.SetBaseOffer(suvWithTurbo)

	suvFullUpgrade := &decorator.PerformanceTires{}
	suvFullUpgrade.SetBaseOffer(suvWithTurboAndAwd)

	fmt.Printf("Base SUV price is %v\nwith turbo it will be %v\nwith all upgrades (AWD and performance tires) it will be %v\n", suv.GetOffer(), suvWithTurbo.GetOffer(), suvFullUpgrade.GetOffer())
}

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
	runDecoratorExample()
}
