package gocollections

import (
	"testing"
)

func TestAll(t *testing.T) {
	cases := []struct {
		enumerable []int
		fun        func(int) bool
		expected   bool
	}{
		{[]int{1, 2, 3, 4, 5}, func(item int) bool { return item > 0 }, true},
		{[]int{1, 2, 3, 4, 5}, func(item int) bool { return item%2 == 0 }, false},
	}

	for _, c := range cases {
		got := All(c.enumerable, c.fun)
		if got != c.expected {
			t.Errorf("TestAll(%#v) == %#v, want %#v", c.enumerable, got, c.expected)
		}
	}
}
