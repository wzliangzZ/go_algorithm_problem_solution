'''
    函数         递归前             回溯时
    build()                       push_up()
    query()     push_down()
   modify()     push_down()       push_up()
'''

class Node:
    def __init__(self, l: int, r: int, v: int = 0, lazy: int = 0):
        self.l = l
        self.r = r
        self.v = v
        self.lazy = lazy


N = 100
#线段树要4倍N
tree = [None for _ in range(4 * N)]

def push_up(u: int):
    tree[u].v = tree[u << 1].v + tree[u << 1 | 1].v          #求区间和
    # tree[u].v = max(tree[u << 1].v, tree[u << 1 | 1].v)     #求区间max
    # tree[u].v = min(tree[u << 1].v, tree[u << 1 | 1].v)     #求区间min

#lazy 标记表示 对以u为根(不包括根)的所有子树生效
def push_down(u: int):
    cur, left, right = tree[u], tree[u << 1], tree[u << 1 | 1]
    if cur.lazy == 0: return
    mid = cur.l + cur.r >> 1

    left.v += (mid - cur.l + 1) * cur.lazy
    right.v += (cur.r - mid) * cur.lazy

    left.lazy += cur.lazy
    right.lazy += cur.lazy
    cur.lazy = 0


#data下标从0开始, 树下标从1开始
#以数组data为值时,建立与之对应值为v的树
def build(u: int, l: int, r: int, data: []) -> None:
    if l == r:
        tree[u] = Node(l, r, data[l - 1])
        return
    tree[u] = Node(l, r)
    mid = l + r >> 1
    build(u << 1, l, mid, data)
    build(u << 1 | 1, mid + 1, r, data)
    push_up(u)


def query(u: int, l: int, r: int) -> int:
    cur = tree[u]
    if l <= cur.l and cur.r <= r: return cur.v

    push_down(u)
    #lv, rv = -0x3f3f3f3f, -0x3f3f3f3f
    lv, rv = 0, 0
    mid = cur.l + cur.r >> 1
    if l <= mid: lv = query(u << 1, l, r)
    if r > mid: rv = query(u << 1 | 1, l, r)
    return lv + rv


def modify(u: int, l: int, r: int, val: int):
    cur = tree[u]
    if l <= cur.l and cur.r <= r:
        cur.v += (cur.r - cur.l + 1) * val
        cur.lazy += val
        return

    push_down(u)
    mid = cur.l + cur.r >> 1
    if l <= mid: modify(u << 1, l, r, val)
    if r > mid: modify(u << 1 | 1, l, r, val)
    push_up(u)
