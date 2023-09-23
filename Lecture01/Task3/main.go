package main

import (
	"reflect"
	"sort"
)

func CompareIntSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	sortedSlice1 := make([]int, len(slice1))
	copy(sortedSlice1, slice1)
	sort.Ints(sortedSlice1)

	sortedSlice2 := make([]int, len(slice2))
	copy(sortedSlice2, slice2)
	sort.Ints(sortedSlice2)
	return reflect.DeepEqual(sortedSlice1, sortedSlice2)
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3, 2, 1}
	result := CompareIntSlices(slice1, slice2)

	if result {
		println("Слайсы совпадают.")
	} else {
		println("Слайсы не совпадают.")
	}
}
