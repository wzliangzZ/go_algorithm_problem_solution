package main

import (
	"bufio"
	. "container/heap"
	"fmt"
)

type pair struct {
	d, cur, bike int64
}

type Heap []pair

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i].d < h[j].d }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *Heap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func CF1915G(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)

	const INF int64 = 1e18
	type node struct{ ne, w int64 }
	s := make([]int64, n)
	g := make([][]node, n)

	for i := 0; i < m; i++ {
		var a, b, c int64
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		g[a] = append(g[a], node{b, c})
		g[b] = append(g[b], node{a, c})
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	dist := make([][]int64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int64, 1001)
		for j := 0; j < 1001; j++ {
			dist[i][j] = INF
		}
	}
	min64 := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	dist[0][1000] = 0
	var h = Heap{pair{0, 0, 1000}}

	for h.Len() > 0 {
		t := Pop(&h).(pair)
		u, k := t.cur, t.bike

		for _, v := range g[u] {
			ne, w := v.ne, v.w
			new_bike := min64(k, s[u])
			if dist[ne][new_bike] > dist[u][k]+w*new_bike {
				dist[ne][new_bike] = dist[u][k] + w*new_bike
				Push(&h, pair{dist[ne][new_bike], ne, new_bike})
			}
		}
	}
	var res int64 = INF
	for _, v := range dist[n-1] {
		res = min64(res, v)
	}
	fmt.Fprintln(out, res)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1915G(in, out)
	}
}
*/
