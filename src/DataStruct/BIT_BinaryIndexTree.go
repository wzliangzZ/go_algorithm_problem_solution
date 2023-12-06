package DataStruct


type Tree []int

func (t Tree) Update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] = max(t[i], val)
	}
}

func (t Tree) Pre_max(i int) int {
	res := -0x3f3f3f3f
	for ; i > 0; i -= i & -i {
		res = max(res, t[i])
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
