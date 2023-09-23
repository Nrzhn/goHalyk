package main

import (
	"fmt"
	"sort"
)

type IntSliceWithSortFunc []int

func (s IntSliceWithSortFunc) Len() int {
	return len(s)
}

func (s IntSliceWithSortFunc) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntSliceWithSortFunc) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSliceWithSortFunc) Sort() {
	sort.Sort(s)
}

func main() {
	intSlice := IntSliceWithSortFunc{5, 2, 9, 1, 5, 6}
	fmt.Println("Исходный слайс:", intSlice)

	intSlice.Sort()
	fmt.Println("Отсортированный слайс:", intSlice)
}
