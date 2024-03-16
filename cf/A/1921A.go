package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1921A(in *bufio.Reader, out *bufio.Writer) {
	arr := make([][]int, 4)
	for i := 0; i < 4; i++ {
		arr[i] = make([]int, 2)
		fmt.Fscan(in, &arr[i][0], &arr[i][1])
	}
	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if abs(arr[i][0]-arr[j][0]) == abs(arr[i][1]-arr[j][1]) {
				fmt.Fprintln(out, abs(arr[i][0]-arr[j][0])*abs(arr[i][1]-arr[j][1]))
				return
			}
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
		CF1921A(in, out)
	}
}
*/
