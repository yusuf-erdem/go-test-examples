package main

import (
	"fmt"
	"testing"
)

//TestReverse test reverse func
func TestReverse(t *testing.T) {
	table := []struct {
		input, output string
	}{
		{"hola", "aloh"},
		{"", ""},
	}
	fmt.Println("**********in test")

	multiply()
	for _, i := range table {
		got := reverse(i.input)
		if i.output != got {
			t.Errorf("Input: %s output: %s correct: %s", i.input, got, i.output)
		}
	}

}

func TestEvenOrSleep(t *testing.T) {
	tt := []struct {
		n        int
		expected error
	}{
		{1, ErrNotEven},
		{3, ErrNotEven},
		{5, ErrNotEven},
		{2, nil},
		{4, nil},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(fmt.Sprint(tc.n), func(t *testing.T) {
			t.Parallel()
			actual := EvenOrSleep(tc.n)

			if tc.expected != actual {
				t.Errorf(`expected "%v", actual "%v"`, tc.expected, actual)
			}
		})
	}
}
