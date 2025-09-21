package gocollections

import (
	"testing"
)

func TestCountWhere(t *testing.T) {
	cases := []struct {
		enumerable []int
		fun        func(int) bool
		expected   int
	}{
		{[]int{1, 2, 3, 4, 5}, func(item int) bool { return true }, 5},
		{[]int{1, 2, 3, 4, 5}, func(item int) bool { return item%2 == 0 }, 2},
		{[]int{}, func(item int) bool { return true }, 0},
	}

	for _, c := range cases {
		got := CountWhere(c.enumerable, c.fun)
		if got != c.expected {
			t.Errorf("TestCountWhere(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}
