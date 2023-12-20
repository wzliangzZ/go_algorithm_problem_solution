package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1521B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	pos := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		if arr[i] < arr[pos] {
			pos = i
		}
	}
	fmt.Fprintln(out, n-1)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for i := 0; i < n; i++ {
		if i != pos {
			fmt.Fprintln(out, pos+1, i+1, arr[pos], arr[pos]+abs(pos-i))
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
		CF1521B(in, out)
	}
}
*/
