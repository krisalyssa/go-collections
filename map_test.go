package gocollections

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestMapIntToInt(t *testing.T) {
	cases := []struct {
		enumerable []int
		fun        func(int) int
		expected   []int
	}{
		{[]int{1, 2, 3, 4, 5}, func(item int) int { return item * 2 }, []int{2, 4, 6, 8, 10}},
	}

	for _, c := range cases {
		got := Map(c.enumerable, c.fun)
		if !slices.Equal(got, c.expected) {
			t.Errorf("TestMapIntToInt(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}

func TestMapStringToInt(t *testing.T) {
	cases := []struct {
		enumerable []string
		fun        func(string) int
		expected   []int
	}{
		{[]string{"1", "2", "3"}, func(item string) int { value, _ := strconv.Atoi(item); return value }, []int{1, 2, 3}},
	}

	for _, c := range cases {
		got := Map(c.enumerable, c.fun)
		if !slices.Equal(got, c.expected) {
			t.Errorf("TestMapStringToInt(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}

func TestMapStringToString(t *testing.T) {
	cases := []struct {
		enumerable []string
		fun        func(string) string
		expected   []string
	}{
		{[]string{"foo", "bar", "baz"}, func(item string) string { return strings.ToUpper(item) }, []string{"FOO", "BAR", "BAZ"}},
	}

	for _, c := range cases {
		got := Map(c.enumerable, c.fun)
		if !slices.Equal(got, c.expected) {
			t.Errorf("TestMapStringToString(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}
