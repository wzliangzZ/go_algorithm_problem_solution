package main

import (
	"bufio"
	"fmt"
	"sort"
	//"os"
)

type SortedList struct {
	Small []int
	Large []int
}

func NewSortedList() SortedList {
	var small []int
	var large []int
	return SortedList{small, large}
}

func (u *SortedList) Add(v int) {
	if len(u.Small) > 6700 {
		u.Large = append(u.Large, u.Small...)
		u.Small = []int{}
		sort.Ints(u.Large)
	}
	if len(u.Small) == 0 || u.Small[len(u.Small)-1] <= v {
		u.Small = append(u.Small, v)
	} else if u.Small[0] >= v {
		u.Small = append([]int{v}, u.Small...)
	} else {
		p := sort.Search(len(u.Small), func(i int) bool { return u.Small[i] >= v })
		t := append([]int{}, u.Small[p:]...) // 实现insert 需要保存临时变量
		u.Small = append(u.Small[:p], v)
		u.Small = append(u.Small, t...)
		//u.Small = append(u.Small[:p], append([]int{v}, u.Small[p:]...)...)
	}
}

func (u SortedList) BisectLeft(v int) int {
	return sort.Search(len(u.Small), func(i int) bool { return u.Small[i] >= v }) + sort.Search(len(u.Large), func(i int) bool { return u.Large[i] >= v })
}

func (u SortedList) Len() int {
	return len(u.Small) + len(u.Large)
}

func CF1915F(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	type pair struct{ a, b int }
	arr := make([]pair, n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		arr[i] = pair{a, b}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].a < arr[j].a
	})
	var res int
	var h = NewSortedList()
	h.Add(arr[0].b)

	for i := 1; i < n; i++ {
		h.Add(arr[i].b)
		idx := h.BisectLeft(arr[i].b)
		res += h.Len() - 1 - idx
	}
	fmt.Fprintln(out, res)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1915F(in, out)
	}
}
*/
