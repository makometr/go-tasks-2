package main

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error") // напишет это, т.к. тип хранящего внутри обхекта не nil, a *customError
		fmt.Printf("%T", err)
		return
	}
	println("ok") // а не это
}

// Для исправления test должен возвращать объекты типа интерфейс error
