package main

import (
	"fmt"

	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/adapter"
	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/bridge"
	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/composite"
	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/decorator"
)

func runDecoratorExample() {
	suv := &decorator.Car{}
	suv.SetPrice(198000)

	suvWithTurbo := &decorator.Turbo{}
	suvWithTurbo.SetBaseOffer(suv)

	suvWithTurboAndAWD := &decorator.AWD{}
	suvWithTurboAndAWD.SetBaseOffer(suvWithTurbo)

	suvFullUpgrade := &decorator.PerformanceTires{}
	suvFullUpgrade.SetBaseOffer(suvWithTurboAndAWD)

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
	aesaRadar := &bridge.AESA{}
	parRadar := &bridge.PAR{}

	f16 := &bridge.Interceptor{}
	f16.SetRadar(aesaRadar)

	b1 := &bridge.Bomber{}
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
