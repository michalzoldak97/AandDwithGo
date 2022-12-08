package plane

import "fmt"

type PAR struct {
}

func (p *PAR) Start() {
	fmt.Println("Starting PAR radar")
}
