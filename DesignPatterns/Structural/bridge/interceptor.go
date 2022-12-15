package bridge

import "fmt"

type Interceptor struct {
	planeRadar Radar
}

func (i *Interceptor) SetRadar(radar Radar) {
	i.planeRadar = radar
}

func (i *Interceptor) StartRadar() {
	fmt.Println("Start interceptor radar")
	i.planeRadar.Start()
}
