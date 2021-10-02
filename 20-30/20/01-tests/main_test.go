package main

import "testing"

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{
			[]int{2, 10, 27, 8},
			47,
		}, {
			[]int{-1, 10, 30, 8},
			47,
		}, {
			[]int{1, 18, 20, 8},
			47,
		}, {
			[]int{0, 11, 27, 8},
			46,
		},
	}

	for _, v := range tests {
		sumOutput := sum(v.data...)
		if sumOutput != v.answer {
			t.Error("Expected value of ", v.answer, " from sum of ", v.data, " but got ", sumOutput)
		}

	}
}
