package main

import (
	"fmt"
)

// Порождающий паттерн
// Способ создания составного объекта
//   - алгоритм создания не зависит от частей и объекта и отношений между ними
//   - обеспечивает различные представления создаваемого объекта2

// Интерфейс строителя объявляет все возможные этапы и шаги конфигурации продукта.
// Все конкретные строители реализуют общий интерфейс по-своему.

type APIparser struct {
	dbConnection string
	parser       string
	logFile      string
}

type APIparserBuilder interface {
	setDBConnection(string)
	setcConfigure([]string)
	setLogFile(string)
	Build() APIparser
}

func main() {
	parsers := make([]APIparser, 0, 2)

	flag := false

	if flag {

		// Эту часть можно выделить и делегировать директору (распорядителю)
		var mgu MGUparserBuilder
		mgu.setDBConnection("k8s.maria-db.shark.us")
		mgu.setcConfigure([]string{"8s", "5s", "3KB", "100"})
		mgu.setLogFile("mgu_log_pasrser.txt")
		parsers = append(parsers, mgu.Build())
	} else {
		// Эту часть можно выделить и делегировать директору (распорядителю)
		var misis MGUparserBuilder
		misis.setDBConnection("k8s.maria-db.shark.us")
		misis.setcConfigure([]string{"8s", "5s", "3KB", "100"})
		misis.setLogFile("mgu_log_pasrser.txt")
		parsers = append(parsers, misis.Build())
	}

	for parser := range parsers {
		fmt.Println(parser)
	}
}
