package main

import "fmt"

// Паттерн команда
//  превращает запросы в объекты со всеми вытекающими возможностями

type command interface {
	execute()
}

type saveDataCommand struct {
	repository interface{}
	maxLife    int
}

func (c saveDataCommand) execute() {
	// if c.maxLife > 5 {
	// 	c.repository.saveStack()
	// } else {
	// 	c.repository.saveHeap()
	// }
	// c.repository.log(c)
}

type pusher interface {
	push()
}

type buttonMenu struct {
	name    string
	styles  map[string]string
	comamnd command
}

func (b buttonMenu) push() {
	fmt.Println("Menu btn pushed!")
	b.comamnd.execute()
}

type buttonCtrlZ struct {
	name    string
	comamnd command
}

func (b buttonCtrlZ) push() {
	fmt.Println("Keyboard btn pushed!")
	b.comamnd.execute()
}

func main() {
	commandToSave := saveDataCommand{repository: nil, maxLife: 5}
	btn1 := buttonMenu{name: "Undo", comamnd: commandToSave}
	btn2 := buttonCtrlZ{name: "Undo", comamnd: commandToSave}

	bb := []pusher{btn1, btn2}

	for _, b := range bb {
		b.push()
	}

}
