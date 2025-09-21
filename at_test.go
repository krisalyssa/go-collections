package gocollections

import (
	"testing"
)

func TestAt(t *testing.T) {
	cases := []struct {
		enumerable []int
		index      int
		expected   int
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 1},
		{[]int{1, 2, 3, 4, 5}, 4, 5},
		{[]int{1, 2, 3, 4, 5}, 5, 5},
		{[]int{1, 2, 3, 4, 5}, -1, 5},
		{[]int{1, 2, 3, 4, 5}, -5, 1},
		{[]int{1, 2, 3, 4, 5}, -6, 1},
	}

	for _, c := range cases {
		got := At(c.enumerable, c.index)
		if got != c.expected {
			t.Errorf("TestAt(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}

}
