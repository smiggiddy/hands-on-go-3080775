// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from error", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	file := os.Args

	// fmt.Println(file[1])
	data, err := os.ReadFile(file[1])
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%T\n", data)
	// convert the bytes to a string
	stringData := string(data)
	// fmt.Printf("%T\n", stringData)

	// initialize a map to store the counts
	counts := map[string]int{"letters": 0, "symbols": 0, "numbers": 0, "words": 0}

	// use the standard library utility package that can help us split the string into words
	stringDataSplit := strings.Split(stringData, " ")

	// capture the length of the words slice
	counts["words"] = len(stringDataSplit)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range stringDataSplit {
		for _, ch := range word {
			if unicode.IsNumber(ch) {
				counts["numbers"]++
			} else if unicode.IsPunct(ch) {
				counts["symbols"]++
			} else {
				counts["letters"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(counts)
}
