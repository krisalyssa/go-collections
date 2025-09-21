package gocollections

import (
	"testing"
)

func TestCount(t *testing.T) {
	cases := []struct {
		enumerable []int
		expected   int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{}, 0},
	}

	for _, c := range cases {
		got := Count(c.enumerable)
		if got != c.expected {
			t.Errorf("TestCount(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}
