package main

import (
	"bufio"
	"fmt"

	//"os"
	"strconv"
	"strings"
)

func P1833(in *bufio.Reader, out *bufio.Writer) {
	var n int
	var s1, s2 string
	var w, v []int
	var tp []bool
	fmt.Fscan(in, &s1, &s2, &n)
	getTime := func() int {
		s1s := strings.Split(s1, ":")
		s2s := strings.Split(s2, ":")
		s2h, _ := strconv.Atoi(s2s[0])
		s2m, _ := strconv.Atoi(s2s[1])
		s1h, _ := strconv.Atoi(s1s[0])
		s1m, _ := strconv.Atoi(s1s[1])
		return (s2h*60 + s2m) - (s1h*60 + s1m)
	}
	t := getTime()
	split := func(a, b, c int) {
		var k int = 1
		for k < c {
			w = append(w, a*k)
			v = append(v, b*k)
			tp = append(tp, false)
			c -= k
			k <<= 1
		}
		w = append(w, a*c)
		v = append(v, b*c)
		tp = append(tp, false)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if c != 0 {
			split(a, b, c)
		} else { //无限选
			w = append(w, a)
			v = append(v, b)
			tp = append(tp, true)
		}
	}
	f := make([]int, t+1)

	for i := 0; i < len(v); i++ {
		var l, r, step int
		if tp[i] {
			l, r, step = w[i], t, 1
		} else {
			l, r, step = t, w[i], -1
		}
		// 乘上 正负一 改变单调性
		for j := l; j*step <= r*step; j += step {
			f[j] = max(f[j], f[j-w[i]]+v[i])
		}
	}
	fmt.Println(f[t])
	//fmt.Fprintln(out, f[t])	// Nothing is compiled: OUTPUT exceeds.
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	//fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		P1833(in, out)
	}
}
*/
