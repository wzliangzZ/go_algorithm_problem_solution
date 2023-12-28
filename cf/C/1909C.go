package main

import (
	"bufio"
	"fmt"
	"sort"
	//"os"
)

func CF1909C(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	type pair struct{ a, b int }
	var arr []pair
	w := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		arr = append(arr, pair{a, 0})
	}
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		arr = append(arr, pair{a, 1})
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i])
	}
	sort.Slice(w, func(i, j int) bool {
		return w[i] > w[j]
	})
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].a < arr[j].a
	})
	var left []int
	var seg []int
	for i := 0; i < len(arr); i++ {
		if arr[i].b == 0 {
			left = append(left, arr[i].a)
		} else {
			seg = append(seg, arr[i].a-left[len(left)-1])
			left = left[:len(left)-1]
		}
	}
	sort.Ints(seg)
	var res int
	for i := 0; i < n; i++ {
		res += seg[i] * w[i]
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
		CF1909C(in, out)
	}
}
*/
