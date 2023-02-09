// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ &/)", want: 1},
		"two":   {input: "812 %6ab//", want: 2},
	}

	lc := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := lc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCounter(t *testing.T) {
	var tests = numberCounter{}
	got := tests.count("hello 5")
	want := 1

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// write a test for symbolCounter.count

func TestSymbolCounter(t *testing.T) {
	got := symbolCounter{}.count("hello!")
	want := 1

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
