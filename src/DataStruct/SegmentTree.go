package DataStruct

/*
    函数         递归前             回溯时
    build()                       push_up()
    query()     push_down()
   modify()     push_down()       push_up()
*/

type Node struct {
	l, r, v, lazy int
}

const MAX_N = 1000000

var tree = [MAX_N]Node{}

func push_up(u int) {
	tree[u].v = tree[u<<1].v + tree[u<<1|1].v //求区间和
	// tree[u].v = max(tree[u << 1].v, tree[u << 1 | 1].v)     //求区间max
	// tree[u].v = min(tree[u << 1].v, tree[u << 1 | 1].v)     //求区间min
}

func push_down(u int) {
	// 获取当前节点和子节点
	cur, left, right := tree[u], tree[u<<1], tree[u<<1|1]

	// 如果当前节点有延迟赋值，则更新子节点和延迟赋值
	if cur.lazy == 0 {
		return
	}

	mid := cur.l + cur.r>>1

	// 更新左子树的值
	left.v += (mid - cur.l + 1) * cur.lazy
	// 更新右子树的值
	right.v += (cur.r - mid) * cur.lazy

	// 更新左子树的延迟赋值
	left.lazy += cur.lazy
	// 更新右子树的延迟赋值
	right.lazy += cur.lazy

	// 清零当前节点的延迟赋值
	cur.lazy = 0
}

//data下标从0开始, 树下标从1开始
func Build(u, l, r int, data []int) {
	// 如果左右子节点范围相同，则将该节点的值初始化为对应数据中的值
	if l == r {
		tree[u] = Node{l, r, data[l-1], 0}
		return
	}

	// 否则，将节点的值初始化为0
	tree[u] = Node{l, r, 0, 0}

	// 计算左右子节点的范围
	mid := (l + r) >> 1

	// 递归构建左右子树
	Build(u<<1, l, mid, data)
	Build(u<<1|1, mid+1, r, data)

	// 更新当前节点的值
	push_up(u)
}

func Query(u int, l int, r int) int {
	// 当前节点为 u
	cur := tree[u]

	// 如果 l 小于等于 cur.l 并且 cur.r 小于等于 r，则直接返回当前节点的值
	if l <= cur.l && cur.r <= r {
		return cur.v
	}

	// 向下推树的标记
	push_down(u)

	// 初始化 lv 和 rv 为 0
	lv, rv := 0, 0
	// 计算中点
	mid := cur.l + cur.r>>1

	// 如果 l 小于等于中点，则递归查询左子树
	if l <= mid {
		lv = Query(u<<1, l, r)
	}

	// 如果 r 大于中点，则递归查询右子树
	if r > mid {
		rv = Query(u<<1|1, l, r)
	}

	// 返回左子树和右子树的查询结果之和
	return lv + rv
}

func Modify(u int, l int, r int, val int) {
	// 当前节点为u对应的节点
	cur := tree[u]

	// 如果查询范围完全包含当前节点的范围
	if l <= cur.l && cur.r <= r {
		// 将当前节点值增加(val*(cur.r-cur.l+1))
		cur.v += (cur.r - cur.l + 1) * val
		// 将当前节点的懒标记增加val
		cur.lazy += val
		return
	}

	// 下推懒标记
	push_down(u)

	// 计算当前节点的中间位置
	mid := cur.l + cur.r>>1

	// 如果查询范围的左边界小于等于当前节点的中间位置
	if l <= mid {
		// 递归修改当前节点的左子树
		Modify(u<<1, l, r, val)
	}

	// 如果查询范围的右边界大于当前节点的中间位置
	if r > mid {
		// 递归修改当前节点的右子树
		Modify(u<<1|1, l, r, val)
	}

	// 上推树结构
	push_up(u)
}
