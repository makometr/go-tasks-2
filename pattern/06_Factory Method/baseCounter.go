package main

// Реализуем интерфейс (как будто базовый абстрактный класс)
type baseViewCounter struct {
	views   int
	viewURL string
}

func (vc *baseViewCounter) addView() {
	vc.views++
}

func (vc *baseViewCounter) reset() {
	vc.views = 0
}

func (vc *baseViewCounter) getTotalViews() int {
	return vc.views
}

func (vc *baseViewCounter) getUrl() string {
	return vc.viewURL
}
