package main

import (
	"bufio"
	"fmt"
	"strconv"
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
