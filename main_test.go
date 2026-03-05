package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Estimate(t *testing.T) {
	cases := []struct {
		levels []int
		total  int
	}{
		{
			levels: nil, // nil slice
			total:  0,
		},
		{
			levels: []int{}, // empty slice
			total:  0,
		},
		{
			levels: []int{2, 0, 1, 0, 1}, // simple test
			total:  2,
		},
		{
			levels: []int{1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, // test 1 from the task
			total:  6,
		},
		{
			levels: []int{3, 0, 2, 0, 4}, // test 2 from the task
			total:  7,
		},
		{
			levels: []int{2, 0, 0, 2}, // test 3 from the task
			total:  4,
		},
		{
			levels: []int{33, 33, 5, 7}, // test 4 from the task
			total:  2,
		},
		{
			levels: []int{5, 2, 7, 3, 4},
			total:  4,
		},
		{
			levels: []int{0, 1, 2, 3, 4, 5, 3, 0, 1},
			total:  1,
		},
		{
			levels: []int{0, 3, 5, 7, 99, 1, 0}, // no space to store oil
			total:  0,
		},
		{
			levels: []int{0, -2, 3, 7, 2, 3}, // test with negative levels
			total:  3,
		},
	}

	for _, tc := range cases {
		if estimate(tc.levels) != tc.total {
			assert.Equal(t, tc.total, estimate(tc.levels))
		}
	}
}
