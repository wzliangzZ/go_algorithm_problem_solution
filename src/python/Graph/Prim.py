import heapq
from collections import defaultdict, deque


dfd = defaultdict(list)     #邻接表
dq = deque()                #双端队列
N, M = 100, 100             #图中点和边数
st = []                     #图中点是否访问过
INF = 0x3f3f3f3f


# 朴素版prim算法 时间复杂度是 O(n^2+m)
st, dist, g = [], [], []
#dist表示 点 到 已确定集合的最短距离
n = 10
def prim() -> int:
    res = 0
    for i in range(n):
        t = -1
        for j in range(n):
            if not st[j] and (t == -1 or dist[t] > dist[j]): t = j
        if i and dist[t] == INF:    #如果不是第一个点，则不存在
            return INF
        if i : res += dist[t]       #除了第一个点，距离都累加上
        st[t] = True
        for j in range(n):
            dist[j] = min(dist[j], g[t][j])
    return res
