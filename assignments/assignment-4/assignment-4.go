package main

import (
	"cmp"
	"fmt"
	"slices"
)

var fp = fmt.Printf

func SortSliceAscending(arr []int) []int {
	duplicateSlice := make([]int, len(arr))
	copy(duplicateSlice, arr)
	slices.Sort(duplicateSlice)
	return duplicateSlice
}

func SortArrayDescending(arr []int) []int {
	duplicateSlice := make([]int, len(arr))
	copy(duplicateSlice, arr)
	sortFunc := func(a, b int) int {
		return cmp.Compare(b, a)
	}
	slices.SortFunc(duplicateSlice, sortFunc)
	return duplicateSlice
}

func SortToOddsAndEvens(arr []int) ([]int, []int) {
	var evens []int
	var odds []int
	for _, v := range arr {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	slices.Sort(odds)
	return odds, SortArrayDescending(evens)
}

func main() {
	intSlice := []int{10, 3, 5, 1, 7, 6, 8, 2, 9, 4}

	fp("original array: %v\n", intSlice)

	fp("the array in ascending order: %v\n", SortSliceAscending(intSlice))

	fp("the array in descending order: %v\n", SortArrayDescending(intSlice))

	odds, evens := SortToOddsAndEvens(intSlice)

	fp("odds in ascending order: %v\n", odds)

	fp("evens in descending order: %v\n", evens)

	fp("the original array again without mutation: %v\n", intSlice)

}
