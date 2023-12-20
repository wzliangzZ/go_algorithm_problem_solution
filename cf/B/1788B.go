package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1788B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var k1, k2, x, y, k, base int
	base = 1
	for n > 0 {
		num := n % 10
		k = num / 2
		if num%2 == 0 {
			x += base * k
			y += base * k
			k1 += k
			k2 += k
		} else {
			if k1 >= k2 {
				y += base * (k + 1)
				x += base * k
				k2 += k + 1
				k1 += k

			} else {
				x += base * (k + 1)
				y += base * k
				k1 += k + 1
				k2 += k
			}
		}
		n /= 10
		base *= 10
	}
	fmt.Fprintln(out, x, y)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		CF1788B(in, out)
	}
}
*/
