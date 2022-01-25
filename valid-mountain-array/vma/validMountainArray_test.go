package vma

import "testing"

type vmaTest struct {
	input []int
	valid bool
}

var tests = []vmaTest{
	{[]int{1, 2, 3, 4, 5, 3, 2, 1}, true},
	{[]int{11, 23, 39, 41, 54, 31, 12, 1}, true},
	{[]int{1, 2, 3, 4, 5, 3, 2, 1, 0, 1}, false},
	{[]int{5, 4, 3, 2, 1}, false},
	{[]int{1, 2, 3, 4, 5}, false},
	{[]int{11, 23, 39, 41, 54, 31, 31, 12, 1}, false},
}

func TestValidMountainArray(t *testing.T) {
	for _, test := range tests {
		if valid := ValidMountainArray(test.input); valid != test.valid {
			t.Errorf("ValidMountainArray(%v) = %v, want %v", test.input, valid, test.valid)
		}
	}
}
