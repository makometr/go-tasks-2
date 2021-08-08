package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type parametres struct {
	after      int
	before     int
	toCount    bool
	ignoreCase bool
	invert     bool
	fixed      bool
	withLine   bool

	pattern  string
	filename string
}

var params parametres

func main() {
	params = parseArgsIntoParams()
	fmt.Println(params)
	data, err := readDataFromFile(params.filename)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)

	doGrep(data, params, os.Stdout)

	// data = doSort(data, params)
	// for _, str := range data {
	// 	fmt.Println(str)
	// }

}

func parseArgsIntoParams() parametres {
	linesAfter := flag.Int("A", 0, "Print N lines after")
	linesBefore := flag.Int("B", 0, "Print N lines before")
	linesNear := flag.Int("C", 0, "Print N lines before and after")
	toCount := flag.Bool("c", false, "Only count lines")
	ignore := flag.Bool("i", false, "Ignore case ")
	invert := flag.Bool("v", false, "Invert results")
	fixed := flag.Bool("F", false, "Fixed pattern")
	withLine := flag.Bool("n", false, "Print lines with number")

	flag.Parse()

	params := parametres{
		after:      *linesAfter,
		before:     *linesBefore,
		toCount:    *toCount,
		ignoreCase: *ignore,
		invert:     *invert,
		fixed:      *fixed,

		withLine: *withLine,
	}

	if *linesNear != 0 {
		params.after = *linesNear
		params.before = *linesNear
	}

	if len(flag.Args()) != 2 {
		return params
	}
	params.pattern = flag.Args()[0]
	params.filename = flag.Args()[1]

	return params
}

func readDataFromFile(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return lines, nil
}

func doGrep(data []string, params parametres, out io.Writer) {
	indices := findIndices(data, params)

	if params.toCount {
		fmt.Fprintf(out, "%d\n", len(indices))
		return
	}

	for _, i := range indices {
		fmt.Fprintf(out, data[i])
	}
}

func findIndices(data []string, params parametres) []int {
	var indices []int

	inverter := func(v bool) bool {
		if params.invert {
			return !v
		}
		return v
	}

	checkByRegex := func(checked string) bool {
		var ans bool
		pattern := params.pattern
		if params.ignoreCase {
			pattern = "(?i)" + pattern
		}
		ans, _ = regexp.MatchString(pattern, checked)
		return ans
	}

	checkByRaw := func(checked string) bool {
		var ans bool
		if params.ignoreCase {
			ans = strings.Contains(strings.ToLower(checked), strings.ToLower(params.pattern))
		} else {
			ans = strings.Contains(checked, params.pattern)
		}
		return ans
	}

	var checker func(string) bool
	if params.fixed {
		checker = checkByRaw
	} else {
		checker = checkByRegex
	}

	for i := 0; i < len(data); i++ {
		if inverter(checker(data[i])) {
			indices = append(indices, i)
		}
	}

	return indices
}
