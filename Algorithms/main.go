package main

import (
	"fmt"

	"github.com/michalzoldak97/AandDwithGo/Algorithms/minmax"
)

func runFindCPU() {

	nums := []int{-35, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, -34, 4, 1029438}

	min, max := minmax.FindNaiveCPU(nums)
	fmt.Print("Min: ", min, " max ", max)
}

func main() {
	runFindCPU()
}
