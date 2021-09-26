package main

import "fmt"

type cheezeSlicer struct {
	cheezeName string
}

func (cs cheezeSlicer) cutcutcut() {
	fmt.Println("Cut cheese ", cs.cheezeName)
}

type meetCooker struct {
	chicken float32
	beaf    float32
}

func (cs meetCooker) prepare() {
	fmt.Printf("Meet prepare:\n\tChicken: %f.2%%\n\tBeaf: %.2f%%\n", cs.chicken, cs.beaf)
}

type doughMaker struct {
	radius int
}

func (dm doughMaker) make() {
	fmt.Println("Prepare the dough with radius", dm.radius)
}
