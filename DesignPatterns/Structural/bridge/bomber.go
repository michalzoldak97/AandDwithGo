package bridge

import "fmt"

type Bomber struct {
	planeRadar Radar
}

func (i *Bomber) SetRadar(radar Radar) {
	i.planeRadar = radar
}

func (i *Bomber) StartRadar() {
	fmt.Println("Start bomber radar")
	i.planeRadar.Start()
}
