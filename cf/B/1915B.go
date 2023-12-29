package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1915B(in *bufio.Reader, out *bufio.Writer) {
	arr := make([]string, 3)
	var x, y int
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &arr[i])
		for j := 0; j < 3; j++ {
			if arr[i][j] == '?' {
				x, y = i, j
			}
		}
	}
	cnt := make([]int, 3)
	for i := 0; i < 3; i++ {
		if arr[x][i] != '?' {
			cnt[arr[x][i]-'A']++
		}
		if arr[i][y] != '?' {
			cnt[arr[i][y]-'A']++
		}
	}
	for i := 0; i < 3; i++ {
		if cnt[i] == 0 {
			fmt.Fprintln(out, string('A'+i))
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
		CF1915B(in, out)
	}
}
*/
