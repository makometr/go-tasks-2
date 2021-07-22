package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s) // 3 2 3
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}

// Копируется по значению, внутренний буфер как указатель, буфер меняется при добавлении нового элемента, а старый нет.
