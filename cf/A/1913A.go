package main

import (
	"bufio"
	"fmt"
	"strconv"
	//"os"
)

func CF1913A(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscan(in, &s)
	var i int
	for i < len(s) {
		if i+1 != len(s) && s[i+1] == '0' {
			i++
		} else {
			break
		}
	}
	a, _ := strconv.Atoi(s[:i+1])
	b, _ := strconv.Atoi(s[i+1:])
	if a >= b {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, a, b)
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
		CF1913A(in, out)
	}

}
*/
