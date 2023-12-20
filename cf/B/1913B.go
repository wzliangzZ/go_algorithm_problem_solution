package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1913B(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	var cnt int
	for _, v := range s {
		if v == '1' {
			cnt++
		}
	}
	f := func(cnt int, c rune) int {
		for i, v := range s {
			if v == c {
				if cnt == 0 {
					return n - i
				}
				cnt--
			}
		}
		return 0
	}
	if cnt == n-cnt {
		fmt.Fprintln(out, 0)
	} else {
		if cnt > n-cnt {
			fmt.Fprintln(out, f(n-cnt, '1'))
		} else {
			fmt.Fprintln(out, f(cnt, '0'))
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
		CF1913B(in, out)
	}

}
*/
