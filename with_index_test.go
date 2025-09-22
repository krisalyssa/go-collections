package gocollections

import (
	"slices"
	"testing"
)

func TestWithIndexInt(t *testing.T) {
	cases := []struct {
		enumerable []int
		expected   []Pair[int, int]
	}{
		{[]int{5, 4, 3, 2, 1}, []Pair[int, int]{{5, 0}, {4, 1}, {3, 2}, {2, 3}, {1, 4}}},
	}

	for _, c := range cases {
		got := WithIndex(c.enumerable)
		if !slices.Equal(got, c.expected) {
			t.Errorf("TestWithIndexInt(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}

func TestWithIndexString(t *testing.T) {
	cases := []struct {
		enumerable []string
		expected   []Pair[string, int]
	}{
		{[]string{"foo", "bar", "baz"}, []Pair[string, int]{{"foo", 0}, {"bar", 1}, {"baz", 2}}},
	}

	for _, c := range cases {
		got := WithIndex(c.enumerable)
		if !slices.Equal(got, c.expected) {
			t.Errorf("TestWithIndexString(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}
