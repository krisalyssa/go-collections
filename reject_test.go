package gocollections

import (
	"slices"
	"strings"
	"testing"
)

func TestRejectInt(t *testing.T) {
	cases := []struct {
		enumerable []int
		fun        func(int) bool
		expected   []int
	}{
		{[]int{1, 2, 3, 4, 5}, func(item int) bool { return item%2 == 0 }, []int{1, 3, 5}},
		{[]int{4, 21, 23, 904}, func(item int) bool { return item > 1000 }, []int{4, 21, 23, 904}},
	}

	for _, c := range cases {
		got := Reject(c.enumerable, c.fun)
		if !slices.Equal(got, c.expected) {
			t.Errorf("TestRejectInt(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}

func TestRejectString(t *testing.T) {
	cases := []struct {
		enumerable []string
		fun        func(string) bool
		expected   []string
	}{
		{[]string{"apple", "pear", "banana"}, func(item string) bool { return strings.Contains(item, "a") }, []string{}},
	}

	for _, c := range cases {
		got := Reject(c.enumerable, c.fun)
		if !slices.Equal(got, c.expected) {
			t.Errorf("TestRejectString(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}
