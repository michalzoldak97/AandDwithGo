package adapter

type DVIAdapter struct {
	DviDisplay *DVIDisplay
}

func (d *DVIAdapter) PlugInDisplayPortCable() {
	d.DviDisplay.PlugInDigitalInputCable()
}
