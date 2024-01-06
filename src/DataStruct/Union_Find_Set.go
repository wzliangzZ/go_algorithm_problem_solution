package DataStruct

const N int = 1000000

var (
	p    = make([]int, N)
	size = make([]int, N)
)

func init() {
	for i := 0; i < N; i++ {
		p[i] = i
		size[i] = 1
	}
}

func Find(x int) int {
	if p[x] != x {
		p[x] = Find(p[x])	// 路径压缩
	}
	return p[x]
}


// 将 x 并入 y
func Union(x, y int) {
	p[Find(x)] = Find(y)

	/*
	启发式合并：将结点数小的合并到结点数大的结点上
	*/
	px, py := Find(x), Find(y)
	if px == py { return }

	// 总认为 x 结点少, 即将 x 并入 y
	if size[px] > size[py] {
		px, py = py, px
	}
	p[px] = py
	size[py] += size[px]
}

func Remove(x int) {
	size[Find(x)]--
	p[x] = x
}

// 将 x 移动到 y
func Move(x, y int) {
    px, py := Find(x), Find(y)
    if px == py {
        return
    }
    p[x] = py
    size[px]--
    size[py]++
}
