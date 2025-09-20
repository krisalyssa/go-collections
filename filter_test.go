package gocollections

import (
	"slices"
	"strings"
	"testing"
)

func TestFilterInt(t *testing.T) {
	cases := []struct {
		enumerable []int
		fun        func(int) bool
		expected   []int
	}{
		{[]int{1, 2, 3, 4, 5}, func(item int) bool { return item%2 == 0 }, []int{2, 4}},
		{[]int{4, 21, 23, 904}, func(item int) bool { return item > 1000 }, []int{}},
	}

	for _, c := range cases {
		got := Filter(c.enumerable, c.fun)
		if !slices.Equal(got, c.expected) {
			t.Errorf("TestFilterInt(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}

func TestFilterString(t *testing.T) {
	cases := []struct {
		enumerable []string
		fun        func(string) bool
		expected   []string
	}{
		{[]string{"apple", "pear", "banana"}, func(item string) bool { return strings.Contains(item, "a") }, []string{"apple", "pear", "banana"}},
	}

	for _, c := range cases {
		got := Filter(c.enumerable, c.fun)
		if !slices.Equal(got, c.expected) {
			t.Errorf("TestFilterString(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}
