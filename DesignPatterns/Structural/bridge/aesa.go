package bridge

import "fmt"

type AESA struct {
}

func (a *AESA) Start() {
	fmt.Println("Starting AESA radar")
}
