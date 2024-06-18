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

class Fenwick:
    __slots__ = 'f'

    def __init__(self, n: int):
        self.f = [0] * n

    def update(self, i: int, val: int) -> None:
        while i < len(self.f):
            self.f[i] += val
            i += i & -i

    def pre(self, i: int) -> int:
        res = 0
        while i > 0:
            res += self.f[i]
            i &= i - 1
        return res

    def query(self, l: int, r: int) -> int:
        if r < l:
            return 0
        return self.pre(r) - self.pre(l - 1)

# L = 100
# tree = [0] * L
# lowbit = lambda x: x & -x
# def update(pos: int = 0, val: int = 0) -> None:
#     while pos < L:
#         tree[pos] += val
#         pos += lowbit(pos)
# def prefix(n: int = 0) -> int:
#     res = 0
#     while n:
#         res += tree[n]
#         n -= lowbit(n)
#     return res

# def query(l: int, r: int) -> int:
#     if r < l:
#         return 0
#     return prefix(r) - prefix(l - 1)
