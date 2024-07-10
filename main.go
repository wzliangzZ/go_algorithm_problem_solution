package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
 
	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		
	}
}

