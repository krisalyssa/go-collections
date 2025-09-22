package gocollections

import (
	"testing"
)

func TestReduceWhileInt(t *testing.T) {
	cases := []struct {
		enumerable []int
		fun        func(int, int) (int, bool)
		expected   int
	}{
		{[]int{1, 2, 3, 4, 5}, func(item int, accumulator int) (int, bool) {
			if item <= 3 {
				return accumulator + item, true
			} else {
				return accumulator, false
			}
		}, 6},
	}

	for _, c := range cases {
		got := ReduceWhile(c.enumerable, 0, c.fun)
		if got != c.expected {
			t.Errorf("TestReduceWhileInt(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}

func TestReduceWhileString(t *testing.T) {
	cases := []struct {
		enumerable []string
		fun        func(string, string) (string, bool)
		expected   string
	}{
		{[]string{"apple", "pear", "banana"}, func(item string, accumulator string) (string, bool) {
			if len(item) >= 5 {
				return accumulator + item, true
			} else {
				return accumulator, false
			}
		}, "apple"},
		{[]string{"apple", "pear", "banana"}, func(item string, accumulator string) (string, bool) {
			if len(item) <= 5 {
				return accumulator + item, true
			} else {
				return accumulator, false
			}
		}, "applepear"},
	}

	for _, c := range cases {
		got := ReduceWhile(c.enumerable, "", c.fun)
		if got != c.expected {
			t.Errorf("TestReduceWhileString(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}

func TestReduceWhileNewValueOnHalt(t *testing.T) {
	cases := []struct {
		enumerable []int
		fun        func(int, int) (int, bool)
		expected   int
	}{
		{[]int{1, 2, 3, 4, 5}, func(item int, accumulator int) (int, bool) {
			if item <= 3 {
				return accumulator + item, true
			} else {
				return item, false
			}
		}, 4},
	}

	for _, c := range cases {
		got := ReduceWhile(c.enumerable, 0, c.fun)
		if got != c.expected {
			t.Errorf("TestReduceWhileNewValueOnHalt(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}
