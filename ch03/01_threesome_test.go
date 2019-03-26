package ch03

import (
	"testing"
)

func TestNewNthStack(t *testing.T) {
	n := 5
	nth := NewNthStack(n)

	if n != nth.n {
		t.Errorf("N expected to be %d, got %d\n", n, nth.n)
	}

	for i := 0; i < n; i++ {
		if i != nth.start[i] {
			t.Errorf("%dth start position expected to be %d, got %d\n", (i + 1), i, nth.start[i])
		}

		if 0 != nth.len[i] {
			t.Errorf("%dth length expected to be zero, got %d\n", (i + 1), nth.len[i])
		}

		if 1 != nth.cap[i] {
			t.Errorf("%dth capacity expected to be 1, got %d\n", (i + 1), nth.cap[i])
		}
	}
}

func TestNthPop(t *testing.T) {
	n, m := 3, 2
	nth := NewNthStack(n)

	nth.storage = make([]int, n*(m+1))
	nth.start = make([]int, n)
	nth.len = make([]int, n)
	nth.cap = make([]int, n)

	for i := 0; i < n; i++ {
		start := i * (m + 1)
		stop := start + m

		nth.start[i] = start
		nth.len[i] = 0
		nth.cap[i] = 1

		for j := start; j < stop; j++ {
			nth.storage[j] = j
			nth.len[i]++
			nth.cap[i]++
		}
	}

	values := [][]int{
		{0, 1},
		{3, 4},
		{6, 7},
	}

	for i := 0; i < nth.n; i++ {
		for j := 0; nth.len[i] > 0; j++ {
			val := values[i][j]

			num, err := nth.Pop(i)
			if err != nil {
				t.Errorf("%dth stack, %dth Pop() expected to be valid, got error %s\n", (i + 1), (j + 1), err)
			}

			if num != val {
				t.Errorf("%dth stack, %dth Pop() expected to be %d, got %d\n", (i + 1), (j + 1), val, num)
			}
		}
	}
}

func TestNthPush(t *testing.T) {
	// TBD
}

func TestNthPeek(t *testing.T) {
	// TBD
}

func TestNthIsEmpty(t *testing.T) {
	// TBD
}
