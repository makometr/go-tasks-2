package main

import (
	"fmt"
	"unsafe"
)

type dataStorage interface {
	load()
	accept(DataStorageVisitor)
}

type stack struct {
	data []byte
}

func (s *stack) load() {
	fmt.Println("Load data from stack")
}

type heap struct {
	segments []*unsafe.Pointer
}

func (s *heap) load() {
	fmt.Println("Load data from heap")
}

// Есть набор связанных структур heap и stack
// Хотим добавить новый функционал со стороны, но не можем руками добавдять новые методы каждый раз.
type DataStorageVisitor interface {
	vistStack(stack)
	visitHeap(heap)
}

func (s *stack) accept(v DataStorageVisitor) {
	v.vistStack(*s)
}

func (h *heap) accept(v DataStorageVisitor) {
	v.visitHeap(*h)
}

// Далее для созданного интерфейса посетитетля реализуем нужную функционональность

type DSVisitorMarshall struct{}

func (dsvm DSVisitorMarshall) vistStack(s stack) {
	fmt.Println("Marshak stack of size", len(s.data))
}

func (dsvm DSVisitorMarshall) visitHeap(h heap) {
	var size int
	for i := 0; i < len(h.segments); i++ {
		// size += h.segments[i]
		size += 1
	}
	fmt.Println("Marshak heap of size", size)
}

// тестирование:

func main() {
	var storages []dataStorage = []dataStorage{
		&heap{segments: []*unsafe.Pointer{nil, nil, nil}},
		&stack{[]byte{1, 2, 3}},
	}

	// Применяем один
	for _, strg := range storages {
		strg.accept(DSVisitorMarshall{})
	}

	// Применяем второй
	for _, strg := range storages {
		strg.accept(DSVisitorToJSON{})
	}
}

type DSVisitorToJSON struct{}

func (dsvm DSVisitorToJSON) vistStack(s stack) {
	fmt.Println("stack jsonify!")
}

func (dsvm DSVisitorToJSON) visitHeap(h heap) {
	fmt.Println("heap jsonify!")
}
