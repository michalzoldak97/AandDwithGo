package main

import (
	"fmt"
	"log"

	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/adapter"
	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/bridge"
	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/composite"
	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/decorator"
	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/facade"
	"github.com/michalzoldak97/AandDwithGo/DesignPatterns/Structural/proxy"
)

func runProxyExample() {
	dbConn := proxy.NewDBServer(111)
	all := 0
	for i := 0; i < 10; i++ {
		res, err := dbConn.RunQuery("SELECT")
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}

		all += res
	}

	fmt.Printf("Finished with %v\n", all)
}

func runFacadeExample() {
	wm, err := facade.NewWashingMachine()
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	wm.StartWashingMachine()
	wm.StopWashingMachine()
}

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
	runFacadeExample()
	runProxyExample()
}
