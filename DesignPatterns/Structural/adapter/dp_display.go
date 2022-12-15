package adapter

import "fmt"

type DPDisplay struct {
}

func (d *DPDisplay) PlugInDisplayPortCable() {
	fmt.Println("Display port cable connected")
}
