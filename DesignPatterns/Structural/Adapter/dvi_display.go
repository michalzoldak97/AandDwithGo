package adapter

import "fmt"

type DVIDisplay struct {
}

func (d *DVIDisplay) PlugInDigitalInputCable() {
	fmt.Println("DVI cable connected")
}
