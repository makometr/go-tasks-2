package main

import "fmt"

// Фасад реализует интерфейс для работы со сложной подсистемой, содержащей много компонентов.

// Фасад создаёт упрощённый интерфейс к подсистеме, не внося в неё никакой добавочной функциональности
// Сама подсистема не знает о существовании Фасада. Классы подсистемы общаются друг с другом напрямую.

// Также можно использовать для разложения подсистемы на отдельные слои.
// Фасад помогает

var cheezeStandartName = "Parmezan"
var chickenSize = 0.6
var beafSize = 0.4
var stRadius = 10

type makePizzaFacade struct {
	cheezer *cheezeSlicer
	meeter  *meetCooker
	dougher *doughMaker
}

func newPizzaFacade() *makePizzaFacade {
	return &makePizzaFacade{
		cheezer: &cheezeSlicer{cheezeStandartName},
		meeter:  &meetCooker{chicken: float32(chickenSize), beaf: float32(beafSize)},
		dougher: &doughMaker{radius: stRadius},
	}
}

func (pf *makePizzaFacade) iDontKnowHowYouCookPizzButMakeItPlease() string {
	fmt.Println("\nPizza making...")
	pf.dougher.make()
	pf.meeter.prepare()
	pf.cheezer.cutcutcut()
	fmt.Printf("Pizza done!\n\n")
	return "Piiza with parmezano!"
}

func main() {
	pizzaMaker := newPizzaFacade()
	pizza := pizzaMaker.iDontKnowHowYouCookPizzButMakeItPlease()

	fmt.Println(pizza)
}
