package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type parametres struct {
	fields        []int
	delim         string
	withDelimOnly bool

	filename string
}

var params parametres

func main() {
	params, err := parseArgsIntoParams()
	if err != nil {
		fmt.Println("Invalid arg!", err)
		return
	}

	fmt.Println(params)

	// data, err := readDataFromFile(params.filename)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// doGrep(data, params, os.Stdout)
}

func parseArgsIntoParams() (parametres, error) {
	fields := flag.String("f", "", "Choose fileds")
	delim := flag.String("d", "", "Choose delim")
	isSep := flag.Bool("s", false, "Separated parametr")

	flag.Parse()

	var columns []int
	for _, column := range strings.Split(*fields, " ") {
		num, err := strconv.Atoi(column)
		if err != nil {
			return parametres{}, err
		}
		columns = append(columns, num)
	}

	if len(*delim) != 1 {
		return parametres{}, fmt.Errorf("Delim should be with size 1")
	}
	if len(flag.Args()) != 1 {
		return parametres{}, fmt.Errorf("Filename should be provided")
	}

	params := parametres{
		fields:        columns,
		delim:         *delim,
		withDelimOnly: *isSep,
		filename:      flag.Args()[0],
	}

	return params, nil
}

// func readDataFromFile(name string) ([]string, error) {
// 	file, err := os.Open(name)
// 	if err != nil {
// 		return []string{}, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	lines := []string{}
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return []string{}, err
// 	}
// 	return lines, nil
// }

// func doGrep(data []string, params parametres, out io.Writer) {
// 	indices := findIndices(data, params)

// 	if params.toCount {
// 		fmt.Fprintf(out, "%d\n", len(indices))
// 		return
// 	}

// 	indices = addAreaIndices(indices, params, len(data))

// 	for i, index := range indices {
// 		if i != len(indices)-1 {
// 			fmt.Fprintf(out, "%s\n", data[index])
// 		} else {
// 			fmt.Fprintf(out, "%s", data[index])
// 		}
// 	}
// }

// func findIndices(data []string, params parametres) []int {
// 	var indices []int

// 	inverter := func(v bool) bool {
// 		if params.invert {
// 			return !v
// 		}
// 		return v
// 	}

// 	checkByRegex := func(checked string) bool {
// 		var ans bool
// 		pattern := params.pattern
// 		if params.ignoreCase {
// 			pattern = "(?i)" + pattern
// 		}
// 		ans, _ = regexp.MatchString(pattern, checked)
// 		return ans
// 	}

// 	checkByRaw := func(checked string) bool {
// 		var ans bool
// 		if params.ignoreCase {
// 			ans = strings.Contains(strings.ToLower(checked), strings.ToLower(params.pattern))
// 		} else {
// 			ans = strings.Contains(checked, params.pattern)
// 		}
// 		return ans
// 	}

// 	var checker func(string) bool
// 	if params.fixed {
// 		checker = checkByRaw
// 	} else {
// 		checker = checkByRegex
// 	}

// 	for i := 0; i < len(data); i++ {
// 		if inverter(checker(data[i])) {
// 			indices = append(indices, i)
// 		}
// 	}

// 	return indices
// }

// func addAreaIndices(inds []int, params parametres, maxIndex int) []int {
// 	if params.after == 0 && params.before == 0 {
// 		return inds
// 	}

// 	unique := make(map[int]struct{})
// 	for _, index := range inds {
// 		unique[index] = struct{}{}

// 		for i := 1; i <= params.before; i++ {
// 			newIndex := index - i
// 			if newIndex >= 0 {
// 				unique[newIndex] = struct{}{}
// 			}
// 		}

// 		for i := 1; i <= params.after; i++ {
// 			newIndex := index + i
// 			if newIndex < maxIndex {
// 				unique[newIndex] = struct{}{}
// 			}
// 		}
// 	}

// 	newIndices := make([]int, 0, len(unique))
// 	for key := range unique {
// 		newIndices = append(newIndices, key)
// 	}
// 	sort.Ints(newIndices)

// 	return newIndices
// }
