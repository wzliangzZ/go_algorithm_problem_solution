#树状数组（Binary Index Tree, BIT）

#树状数组区间求和问题
# 1. 维护原数组：     单点修改/区间查询   时间复杂度：O(logn)
# 2. 维护差分数组：   区间修改/单点查询   时间复杂度：O(logn)
# 3. 同时维护 差分数组 和 i*差分数组      区间修改/区间查询   时间复杂度：O(logn)
#树状数组（Binary Index Tree, BIT）

# 1. 维护原数组：     单点修改/区间查询   时间复杂度：O(logn)
# 2. 维护差分数组：   区间修改/单点查询   时间复杂度：O(logn)
# 3. 同时维护 差分数组 和 i*差分数组      区间修改/区间查询   时间复杂度：O(logn)

#tree[n]维护的是前n项和，以lowbit(n)为划分依据的若干区间的总和
L = 100
tree = [0] * L


lowbit = lambda x: x & -x

def update(pos: int = 0, val: int = 0) -> None:
    while pos < L:
        tree[pos] += val
        pos += lowbit(pos)

def prefix(n: int = 0) -> int:
    res = 0
    while n:
        res += tree[n]
        n -= lowbit(n)

query = lambda l, r: prefix(r) - prefix(l - 1)