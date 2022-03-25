package usort

import "sort"

type Slice[T any] []T

type SortSlice[T any] struct {
	Slice[T]
	less func(i, j int) bool
}

func USort[T any](sli []T, sortFunc func(i, j int) bool) {
	ss := SortSlice[T]{Slice: sli}
	ss.less = sortFunc
	sort.Sort(ss)
}

func (s SortSlice[T]) Len() int {
	return len(s.Slice)
}

func (s SortSlice[T]) Swap(i, j int) {
	(s.Slice)[i], (s.Slice)[j] = (s.Slice)[j], (s.Slice)[i]
}

func (s SortSlice[T]) Less(i, j int) bool {
	return s.less(i, j)
}
