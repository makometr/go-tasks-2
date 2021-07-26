package main

// Порождающий паттерн
// определяет общий интерфейс для создания объектов в базовом классе
// А подклассы изменяют тип создаваемых объектов.

// способ скрыть логику создания создаваемых экземпляров.

// определяет все методы, которыми должно обладать cчётчик просмотров.
// Т.е. теперь объёкты, реализующие интерфейс связаны не только общим набором методов, но и общим встроенным объектом
type viewCounter interface {
	addView()
	reset()
	getTotalViews() int
	getUrl() string
}

func getCounterView(viewType string, url string) viewCounter {
	switch viewType {
	case "vk":
		return newVKCounter(url)
	case "youtube":
		return newYoutubeCounter(url)
	}
	return nil
}

func main() {

}
