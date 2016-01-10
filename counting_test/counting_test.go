package counting_test

import (
	"testing"

	"github.com/umahmood/counting"
)

var tests = []struct {
	input    []int
	expected []int
}{
	{
		input:    []int{3, 6, 4, 1, 3, 4, 1, 4},
		expected: []int{1, 1, 3, 3, 4, 4, 4, 6},
	},
	{
		input:    []int{1, 4, 7, 2, 1, 3, 2, 1, 4, 2, 3, 2, 1},
		expected: []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 4, 7},
	},
	{
		input:    []int{6, 2, 1, 1, 0},
		expected: []int{0, 1, 1, 2, 6},
	},
	{
		input:    []int{2048, 3, 1, 5},
		expected: []int{1, 3, 5, 2048},
	},
	{
		input:    []int{42, -88, 1, -3, 2, 5},
		expected: []int{-88, -3, 1, 2, 5, 42},
	},
	{
		input:    []int{-42, -88, -1, -3, 0, -5},
		expected: []int{-88, -42, -5, -3, -1, 0},
	},
	{
		input:    []int{},
		expected: []int{},
	},
}

func TestSort(t *testing.T) {
	// compareIntSlices checks if two integer slices contain elements in the
	// same order.
	compareIntSlices := func(a, b []int) bool {
		if len(a) != len(b) || a == nil || b == nil {
			return false
		}

		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	for _, test := range tests {
		got := counting.Sort(test.input)
		if !compareIntSlices(got, test.expected) {
			t.Errorf("Sort fail - got %v expected %v", got, test.expected)
		}
	}
}

func TestSortInputIsNil(t *testing.T) {
	var input []int

	got := counting.Sort(input)

	if got != nil {
		t.Errorf("Sort fail - got %v expected nil", got)
	}
}
