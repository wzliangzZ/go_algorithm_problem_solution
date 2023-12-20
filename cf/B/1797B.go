package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1797B(in *bufio.Reader, out *bufio.Writer) {
	var n, k, cnt int
	fmt.Fscan(in, &n, &k)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &arr[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] != arr[n-1-i][n-1-j] {
				cnt++
			}
		}
	}
	cnt /= 2
	if k < cnt {
		fmt.Fprintln(out, "NO")
	} else if n%2 != 0 || (k-cnt)%2 == 0 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
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
		CF1797B(in, out)
	}
}
*/
