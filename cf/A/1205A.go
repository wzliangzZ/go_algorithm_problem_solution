package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1205A(in *bufio.Reader, out *bufio.Writer) {
	// var T int
	// for fmt.Fscan(in, &T); T > 0; T-- {

	// }
	var n int
	fmt.Fscan(in, &n)
	if n%2 == 0 {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
		arr := make([]int, n)
		var flag bool
		var k int = 1
		for i := 0; i < n; i++ {
			if !flag {
				fmt.Fprint(out, k, " ")
				k++
				arr[i] = k
				k++
			} else {
				arr[i] = k
				k++
				fmt.Fprint(out, k, " ")
				k++
			}
			flag = !flag
		}
		for i := 0; i < n; i++ {
			fmt.Fprint(out, arr[i], " ")
		}
		fmt.Fprintln(out)
	}

}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	CF1205A(in, out)
}
*/
