from collections import defaultdict

"""
并查集
深搜，回时指父，离时搜根
O(n+m)
"""

N = 10 ** 5     #结点数
M = 10 ** 5     #查询数
dfd = defaultdict(list)
#(key, list(tuple()) ) （u,(v,idx), (v, idx))  双向存
query = defaultdict(list)
# 并查集
fa = [i for i in range(N)]

vis, res = [False] * N, [M]


def find(x: int = 0) -> int:
    if fa[x] != x: fa[x] = find(fa[x])
    return fa[x]

def tarjan(u: int = 0) -> None:
    vis[u] = True
    for ne in dfd[u]:
        if not vis[ne]:
            tarjan(ne)
            fa[ne] = u  #回溯时，v指向u

    # 离开u时，枚举LCA
    for v, idx in query[u]:
        if vis[v]: res[idx] = find(v)