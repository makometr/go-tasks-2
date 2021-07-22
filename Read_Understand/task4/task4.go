package main

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// close(ch)
	}()

	for n := range ch {
		println(n)
	}
}

// Программа напечатает числа от 0 до 9. Затем главная горутина зависнет в ожидании закрытия канала ch.
// т.к. конструкция for range channel будет ожидать и читать значения, пока канал открыт.
