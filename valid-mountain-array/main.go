package main

import (
	"fmt"

	"github.com/geekmdtravis/go-leetcode/valid-mountain-array/vma"
)

func main() {
	valid := []int{1, 2, 3, 4, 5, 3, 2, 1}
	invalid := []int{1, 2, 3, 4, 5, 3, 2, 1, 0}
	fmt.Println(vma.ValidMountainArray(valid))
	fmt.Println(vma.ValidMountainArray(invalid))
}
