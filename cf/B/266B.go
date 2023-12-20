package main
 
import (
	"bufio"
	"fmt"
	// "os"
)
func CF266B(in *bufio.Reader, out *bufio.Writer) {
	var n, t int
	fmt.Fscan(in, &n, &t)
	var s string
	fmt.Fscan(in, &s)
	arr := []byte(s)
	for t > 0{
		var n_arr[]byte
		n_arr = append(n_arr, arr[0])
		for i := 1; i < n; i++ {
			if arr[i] == 'G' && arr[i - 1] == 'B'{
				n_arr[len(n_arr) - 1] = 'G'
				n_arr = append(n_arr, 'B')
			}else{
				n_arr = append(n_arr, arr[i])
			}
		}
		arr = n_arr
		t--
	}
	fmt.Fprintln(out, string(arr))
}
/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
 
	var tot int = 1
	//fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		CF266B(in, out)
	}
}
*/