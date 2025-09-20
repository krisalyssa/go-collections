package gocollections

import (
	"testing"
)

func TestReduceInt(t *testing.T) {
	cases := []struct {
		enumerable []int
		fun        func(int, int) int
		expected   int
	}{
		{[]int{1, 2, 3, 4, 5}, func(accumulator int, item int) int { return accumulator + item }, 15},
	}

	for _, c := range cases {
		got := Reduce(c.enumerable, 0, c.fun)
		if got != c.expected {
			t.Errorf("TestReduceInt(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}

func TestReduceString(t *testing.T) {
	cases := []struct {
		enumerable []string
		fun        func(string, string) string
		expected   string
	}{
		{[]string{"apple", "pear", "banana"}, func(item string, accumulator string) string { return accumulator + item }, "peachapplepearbanana"},
	}

	for _, c := range cases {
		got := Reduce(c.enumerable, "peach", c.fun)
		if got != c.expected {
			t.Errorf("TestReduceString(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}
