package main

import "fmt"

// поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков.
// Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

type car struct {
	wheels       []int
	transmission *string
	engine       *string
}

type carReport struct {
	car        *car
	healthRate int
}

func main() {
	var cw checkWheels = checkWheels{&endOfCheck{}}
	var car car = car{wheels: []int{1, 2, 3, 4}}

	report := &carReport{car: &car}
	cw.check(report)

	fmt.Println("Result: ", report.healthRate)
}
