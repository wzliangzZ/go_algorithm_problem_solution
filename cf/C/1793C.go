package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1793C(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	var i, j int = 0, n - 1
	var l, r int = 1, n
	for i < j {
		if l < arr[i] && arr[i] < r && l < arr[j] && arr[j] < r {
			fmt.Fprintln(out, i+1, j+1)
			return
		}
		if arr[i] == l {
			i++
			l++
		} else if arr[i] == r {
			i++
			r--
		}
		if arr[j] == l {
			j--
			l++
		} else if arr[j] == r {
			j--
			r--
		}
	}
	fmt.Fprintln(out, -1)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1793C(in, out)
	}
}
*/
