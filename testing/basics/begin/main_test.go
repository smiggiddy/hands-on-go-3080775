// testing/basics/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

// write a test for sum

func TestSum(t *testing.T) {
	got := sum(1, 2, 3)
	want := 6
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// write a TestMain for setup and teardown
func TestMain(m *testing.M) {
	fmt.Println("Testing setup")
	m.Run()
	fmt.Println("Testing teardown ran")
}
