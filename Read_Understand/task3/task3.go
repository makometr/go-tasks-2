package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}
func main() {
	err := Foo()
	fmt.Println(err) // nil
	fmt.Printf("\n%T - Тип хранимого объекта под капотом\n", err)
	fmt.Printf("%v - Значение этого объекта под капотом\n", err)
	fmt.Println("err == nil:", err == nil) // false

	err = nil
	fmt.Printf("\n%T - Тип хранимого объекта под капотом\n", err)
	fmt.Printf("%v - Значение этого объекта под капотом\n", err)
	fmt.Println("err == nil:", err == nil) // true
}

// error - интерфейс. Внутри реализации он хранит данные и типе хранимого объекта (%T) и указатель на сам хранимый объект.
// Данные о типе: самого интерфейса, хранимой структуры (объекта) и таблицу с методами.
//
// False т.к. сам объект интерфеса - это не nil. Да, он хранит в себе nil, но nil`ом становится, когда и его тип тоже является nil.
// В первом случае тип у хранимого объекта  - есть. А значит объект интерфейса, его содержащий, нилу равен не будет.
