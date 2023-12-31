package DataStruct

import (
	"sort"
)

/*
	有序数组

Add()			O(logn)
Delete()		待实现...
BisectLeft()	O(logn)
*/
type SortedList struct {
	Small []int
	Large []int
}

func NewSortedList() SortedList {
	return SortedList{[]int{}, []int{}}
}

func (sl *SortedList) Add(v int) {
	if len(sl.Small) > 7222 {
		sl.Large = append(sl.Large, sl.Small...)
		sl.Small = []int{}
		sort.Ints(sl.Large)
	}
	if len(sl.Small) == 0 || sl.Small[len(sl.Small)-1] <= v {
		sl.Small = append(sl.Small, v)
	} else if sl.Small[0] >= v {
		sl.Small = append([]int{v}, sl.Small...)
	} else {
		p := sort.Search(len(sl.Small), func(i int) bool { return sl.Small[i] >= v })
		t := append([]int{}, sl.Small[p:]...)
		sl.Small = append(sl.Small[:p], v)
		sl.Small = append(sl.Small, t...)
		//sl.Small = append(sl.Small[:p], append([]int{v}, sl.Small[p:]...)...)
	}
}

// 大于等于V的最左边的数的下标,不存在返回长度
func (sl SortedList) BisectLeft(v int) int {
	return sort.Search(len(sl.Small), func(i int) bool { return sl.Small[i] >= v }) + sort.Search(len(sl.Large), func(i int) bool { return sl.Large[i] >= v })
}

func (sl SortedList) Len() int {
	return len(sl.Small) + len(sl.Large)
}
