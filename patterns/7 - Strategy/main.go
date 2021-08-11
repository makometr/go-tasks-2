package main

import "fmt"

// Стрратегия
// определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс,
// после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

// Объект, над которым будет производиться работа
type bigPieceOfData struct {
	data     [][][][][]interface{}
	sortAlgo dataSorter
}

func (bpd *bigPieceOfData) setSorterAlgo(algo dataSorter) {
	bpd.sortAlgo = algo
}

func (bpd *bigPieceOfData) callSort() {
	fmt.Println("Call sort algo!")
	bpd.sortAlgo.doSort(bpd)
}

// Интерфейс этой самой работы
type dataSorter interface {
	doSort(*bigPieceOfData)
}

// Конкретная реализация типа работ
type quickSort struct{}

func (s quickSort) doSort(bpf *bigPieceOfData) {
	fmt.Println("Sort data with quick algo!")

}

// Конкретная реализация другого типа работ
type mergeSort struct{}

func (s mergeSort) doSort(bpf *bigPieceOfData) {
	fmt.Println("Sort data with merge algo!")
}

func main() {
	var piece bigPieceOfData

	sortData := func() {
		if len(piece.data) > 10000 {
			piece.setSorterAlgo(quickSort{})
		} else {
			piece.setSorterAlgo(mergeSort{})
		}
		piece.callSort()
	}

	sortData()
}
