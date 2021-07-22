package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) *
				time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			// select {
			// case v := <-a:
			// 	c <- v
			// case v := <-b:
			// 	c <- v
			// }
			select {
			case x, ok := <-a:
				if ok {
					c <- x
				} else {
					a = nil
				}
			case x, ok := <-b:
				if ok {
					c <- x
				} else {
					b = nil
				}
			}

			if a == nil && b == nil {
				break
			}
		}
		close(c) // а ещё надо закрыть канал!
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

// Как только канал а или б закроется - он будет немедленно доступен для получения значения из него (ноль).
// Получем, что после получения всех значений проограмма не прекратится, а войдёт в бесконечный цикл, выводящий нули.

// Исправленное решение походит в качестве костыля при малом и известном количестве каналов для слияния.
// Если в merge приходит ...<-chan int, то нужно использовать дополнительные средства синхронизации с wg и синхронизирующим каналом.
// Паттерн - fun in
