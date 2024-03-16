package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1236C(in *bufio.Reader, out *bufio.Writer) {
	// var T int
	// for fmt.Fscan(in, &T); T > 0; T-- {

	// }
	var n int
	fmt.Fscan(in, &n)
	arr := make([][]int, n)
	var k int = 1
	for j := 0; j < n; j++ {
		var s, e, step, c int = 0, n - 1, 1, 1
		if j % 2 != 0 { s, e, step, c = n - 1, 0, -1, -1 }
		for i := s; i * c <= e; i += step {
			if j == 0 {
				arr[i] = make([]int, n)
			}
			arr[i][j] = k
			k++
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(out, arr[i][j], " ")
		}
		fmt.Fprintln(out)
	}
}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	CF1236C(in, out)
// }
