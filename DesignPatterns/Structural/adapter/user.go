package adapter

import "fmt"

type User struct {
}

func (u *User) PlugInDisplayPortIntoDisplay(d Display) {
	d.PlugInDisplayPortCable()
	fmt.Println("Inserting display plug")
}
