package main

import (
	"fmt"
	"unicode"
)

// Паттерн состояние позволяет лучше структурировать внутрениюю логику,
// зависящую от параметром объекта (его состояния)
// При этом различные состояния знают друг о друге и о своем вместителе
// и могут (скорее всего должны) взаимодействовать друг с другом.

// Распознование последовательностей из трёх букв и двух цифр.

type myParser struct {
	curState ParserState
}

func newMyParser() *myParser {
	return &myParser{curState: &parserStateReadThreeWords{}}
}

func (p myParser) parse(str string) (bool, error) {
	for _, r := range str {
		if p.curState.putSymbol(&p, r) == false {
			fmt.Printf("%c incorrect\n", r)
			return false, nil
		}
	}
	return true, nil
}

type ParserState interface {
	putSymbol(*myParser, rune) bool
}

type parserStateReadThreeWords struct {
	wordsRead int
}

func (ps *parserStateReadThreeWords) putSymbol(p *myParser, r rune) bool {
	if !unicode.IsLetter(r) {
		return false
	}
	ps.wordsRead++
	if ps.wordsRead > 3 {
		return false
	}
	if ps.wordsRead == 3 {
		p.curState = &parserStateReadOneDigit{}
	}
	return true
}

type parserStateReadOneDigit struct{}

func (ps *parserStateReadOneDigit) putSymbol(p *myParser, r rune) bool {
	if !unicode.IsDigit(r) {
		return false
	}
	p.curState = &parserStateReadThreeWords{}
	return true
}

func main() {
	{
		stringToParse := []rune("asd3asd4")
		parser := newMyParser()
		res, err := parser.parse(string(stringToParse))
		if err != nil {
			return
		}
		fmt.Println("Result:", res)
	}

	{
		stringToParse := []rune("asd3asd4555")
		parser := newMyParser()
		res, err := parser.parse(string(stringToParse))
		if err != nil {
			return
		}
		fmt.Println("Result:", res)
	}

}
