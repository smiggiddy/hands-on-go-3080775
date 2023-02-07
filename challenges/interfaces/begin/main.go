// challenges/interfaces/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

type counter interface {
	name() string
	count(input string) int
}

type letterCounter struct{ identifier string }

type numberCounter struct{ designation string }

type symbolCounter struct{ label string }

// func (s interface{}) count(i string) int {
// 	var sum int
// 	for _, l := range i {
// 		switch b := s.(type) {
// 		case letterCounter:
// 			sum++
// 		case numberCounter:
// 			if unicode.IsNumber(l) {
// 				sum++
// 			}

// 		case symbolCounter:
// 			if unicode.IsSymbol(l) {
// 				sum++
// 			}

// 		}
// 	}
// 	return sum
// }

func (s symbolCounter) count(input string) (sum int) {

	for _, l := range input {
		if !unicode.IsLetter(l) && !unicode.IsNumber(l) {
			sum++
		}
	}
	return
}

func (n numberCounter) count(input string) (sum int) {

	for _, l := range input {
		if unicode.IsNumber(l) {
			sum++
		}
	}
	return
}

func (let letterCounter) count(input string) (sum int) {

	for _, l := range input {
		if unicode.IsLetter(l) {
			sum++
		}
	}
	return
}

func (s symbolCounter) name() string { return fmt.Sprintf("%v", s.label) }
func (n numberCounter) name() string { return fmt.Sprintf("%v", n.designation) }
func (l letterCounter) name() string { return fmt.Sprintf("%v", l.identifier) }

func doAnalysis(data string, counters ...counter) map[string]int {
	// initialize a map to store the counts
	analysis := make(map[string]int)

	// capture the length of the words in the data
	analysis["words"] = len(strings.Fields(data))

	// loop over the counters and use their name as the key
	for _, c := range counters {
		analysis[c.name()] = c.count(data)
	}

	// return the map
	return analysis
}

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	// use os package to read the file specified as a command line argument
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	data := string(bs)
	spew.Dump(data)

	s := numberCounter{designation: "numbers"}
	l := letterCounter{identifier: "letters"}
	sy := symbolCounter{label: "symbols"}
	// call doAnalysis and pass in the data and the counters
	fmt.Println(s.designation)
	analysis := doAnalysis(data, s, l, sy)
	// dump the map to the console using the spew package
	spew.Dump(analysis)
}
