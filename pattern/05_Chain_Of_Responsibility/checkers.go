package main

import "fmt"

type carChecker interface {
	check(*carReport)
	setNext(carChecker)
}

type checkWheels struct {
	next carChecker
}

func (ch *checkWheels) check(cr *carReport) {
	fmt.Println("Check wheels!")
	if cr.car.wheels != nil {
		cr.healthRate += 2
		ch.setNext(&checkEngine{})
	} else {
		fmt.Println("No wheels!")
		ch.setNext(&endOfCheck{})
	}
	ch.next.check(cr)
}

func (ch *checkWheels) setNext(cc carChecker) {
	ch.next = cc
}

type checkEngine struct {
	next carChecker
}

func (ch *checkEngine) check(cr *carReport) {
	fmt.Println("Check engine!")
	if cr.car.engine != nil {
		cr.healthRate += 4
		if *cr.car.engine == "" {
			ch.setNext(&checkTransmission{&endOfCheck{}})
		} else {
			fmt.Println("Engine ok!")
			cr.healthRate += 5
			ch.setNext(&endOfCheck{})
		}
	} else {
		fmt.Println("No engine!")
		ch.setNext(&endOfCheck{})
	}
	ch.next.check(cr)
}

func (ch *checkEngine) setNext(cc carChecker) {
	ch.next = cc
}

type checkTransmission struct {
	next carChecker
}

func (ch *checkTransmission) check(cr *carReport) {
	fmt.Println("Check transmission!")
	if cr.car.transmission != nil {
		cr.healthRate += 1
	} else {
		fmt.Println("No transmission!")
	}
	ch.next.check(cr)
}

func (ch *checkTransmission) setNext(cc carChecker) {
	ch.next = cc
}

type endOfCheck struct {
}

func (ch *endOfCheck) check(cr *carReport) {
	fmt.Println("End of check!")

}

func (ch *endOfCheck) setNext(cc carChecker) {
}
