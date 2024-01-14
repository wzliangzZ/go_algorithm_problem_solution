package string

// 返回 s 中的所有 t 的起始位置

func KMP(s string, t string) *[]int {
	var n, m int = len(s), len(t)
	next := make([]int, m)
	var res []int

	getNext := func() {
		var cnt = 0
		// 求next从 1 开始
		for i := 1; i < m; i++ {
			for cnt > 0 && t[cnt] != t[i] {
				cnt = next[cnt-1]
			}
			if t[cnt] == t[i] {
				cnt++
			}
			next[i] = cnt
		}
	}

	search := func() {
		var j int
		// 匹配从 0 开始
		for i := 0; i < n; i++ {
			for j > 0 && s[i] != t[j] {
				j = next[j-1]
			}
			if s[i] == t[j] {
				j++
			}
			if j == m {
				res = append(res, i-m+1)
				j = next[j-1]
			}
		}
	}

	getNext()
	search()
	return &res
}
