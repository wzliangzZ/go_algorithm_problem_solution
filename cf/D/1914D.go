package main

import (
	"bufio"
	"fmt"
	"sort"
	//"os"
)

type pair struct {
	v, idx int
}

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Less(i, j int) bool {
	return p[i].v > p[j].v
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func CF1914D(in *bufio.Reader, out *bufio.Writer) {
	var n, t int
	fmt.Fscan(in, &n)
	arra := make(pairs, n)
	arrb := make(pairs, n)
	arrc := make(pairs, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t)
		arra[i] = pair{t, i}
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t)
		arrb[i] = pair{t, i}
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t)
		arrc[i] = pair{t, i}
	}
	sort.Sort(arra)
	sort.Sort(arrb)
	sort.Sort(arrc)
	var res int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if arra[i].idx != arrb[j].idx && arra[i].idx != arrc[k].idx && arrb[j].idx != arrc[k].idx {
					res = max(res, arra[i].v+arrb[j].v+arrc[k].v)
				}
			}
		}
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
		CF1914D(in, out)
	}

}
*/
