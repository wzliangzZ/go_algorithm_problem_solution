package main

import (
	"bufio"
	"fmt"
	"sort"
	//"os"
)

type rank struct {
	score  int
	idx    int
	pently int
}
type ranks []*rank

func (rs ranks) Len() int {
	return len(rs)
}
func (rs ranks) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
func (rs ranks) Less(i, j int) bool {
	if rs[i].score == rs[j].score {
		if rs[i].pently == rs[j].pently {
			return rs[i].idx < rs[j].idx
		}
		return rs[i].pently < rs[j].pently
	} else {
		return rs[i].score > rs[j].score
	}
}
func CF1864C(in *bufio.Reader, out *bufio.Writer) {
	var n, m, h int
	fmt.Fscan(in, &n, &m, &h)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &arr[i][j])
		}
		sort.Ints(arr[i])
	}

	calc := func(idx int, arr []int, h int, res *ranks) {
		var sc, p, cur int
		for i := 0; i < len(arr); i++ {
			if cur+arr[i] > h {
				break
			}
			cur += arr[i]
			sc++
			p += cur
		}
		*res = append(*res, &rank{sc, idx, p})
	}
	var res ranks
	for i := 0; i < n; i++ {
		calc(i, arr[i], h, &res)
	}
	sort.Sort(res)
	for i, rk := range res {
		if rk.idx == 0 {
			fmt.Fprintln(out, i+1)
			return
		}
	}
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		CF1864C(in, out)
	}

}
*/
