package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1854A2(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	var mxi, mni, pos, neg int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		if arr[i] > 0 {
			if arr[i] > arr[mxi] {
				mxi = i
			}
			pos++
		} else if arr[i] < 0 {
			if arr[i] < arr[mni] {
				mni = i
			}
			neg++
		}
	}
	type pair struct{ a, b int }
	var res []pair
	makepos := func() {
		for i := 0; i < n; i++ {
			if arr[i] < 0 {
				res = append(res, pair{i, mxi})
			}
			if i > 0 {
				res = append(res, pair{i, i - 1})
			}
		}
	}
	makeneg := func() {
		for i := n - 1; i >= 0; i-- {
			if arr[i] > 0 {
				res = append(res, pair{i, mni})
			}
			if i < n-1 {
				res = append(res, pair{i, i + 1})
			}
		}
	}

	if arr[mxi] >= -arr[mni] {
		if neg <= 12 {
			makepos()
		} else {
			for i := 0; i < 5; i++ {
				res = append(res, pair{mni, mni})
			}
			makeneg()
		}
	} else {
		if pos <= 12 {
			makeneg()
		} else {
			for i := 0; i < 5; i++ {
				res = append(res, pair{mxi, mxi})
			}
			makepos()
		}
	}
	fmt.Fprintln(out, len(res))
	for _, v := range res {
		fmt.Fprintln(out, v.a+1, v.b+1)
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
		CF1854A2(in, out)
	}
}
*/
