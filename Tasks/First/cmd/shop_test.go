package main

import (
	"fmt"
	"testing"
)

func TestGetSections(t *testing.T) {
	s := newShop()
	for sec := range s.getSections() {
		fmt.Println(sec)
	}
}
