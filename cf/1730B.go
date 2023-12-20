package main

import (
	"bufio"
	"fmt"
)

func CF1730B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arrx := make([]int, n)
	arrt := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arrx[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arrt[i])
	}

	check := func(mid int) (bool, float32) {
		a, b := -1000000000, 1000000000
		for i := 0; i < n; i++ {
			if mid-arrt[i] < 0 {
				return false, -1
			}
			na, nb := arrx[i]-(mid-arrt[i]), arrx[i]+(mid-arrt[i])
			if na > b || nb < a {
				return false, -1
			}
			if na > a {
				a = na
			}
			if nb < b {
				b = nb
			}
		}
		return true, float32(a+b) / 2
	}

	var res float32
	var l, r int = 0, 1000000000

	for l < r {
		mid := (l + r) / 2
		if ok, x := check(mid); ok {
			res = x
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Fprintln(out, res)
}
