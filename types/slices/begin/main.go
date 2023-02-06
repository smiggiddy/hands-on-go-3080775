// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	names := make([]string, 0)

	// append 3 names to the slice
	names = append(names, "smig")
	names = append(names, "testing")
	names = append(names, "simple")

	// print the slice
	fmt.Println(names)

	// slice the slice using syntax slice[low:high]
	// [Jane Mary]
	// [Jane Mary]
	// [John]
	// [John Jane Mary]
}
