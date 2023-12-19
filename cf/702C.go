package main

import (
	"bufio"
	"fmt"
)

func CF702C(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	city := make([]int, n)
	tower := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &city[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &tower[i])
	}

	check := func(t int) bool {
		var cur, i int
		var a, b int = tower[cur] - t, tower[cur] + t
		for i < n && cur < m {
			if city[i] < a {
				return false
			}
			if city[i] > b {
				cur++
				if cur == m {
					break
				}
				a, b = tower[cur]-t, tower[cur]+t
				continue
			}
			i++
		}
		return i == n
	}

	var l, r int = 0, 0x3f3f3f3f * 2
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Fprintln(out, l)
}
