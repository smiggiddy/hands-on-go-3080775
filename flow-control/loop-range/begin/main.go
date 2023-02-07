// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	num := []int{100, 200, 300}

	// use for-range to iterate over the slice and print each value
	for i, n := range num {
		fmt.Println(i, ":", n)
	}

	// declare a map of strings to ints
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	// use for-range to iterate over the map and print each key/value pair
	for key, val := range m {
		fmt.Println(key, val)
	}
}
