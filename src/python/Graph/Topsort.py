import heapq
from collections import defaultdict, deque


dfd = defaultdict(list)     #邻接表
dq = deque()                #双端队列
N, M = 100, 100             #图中点和边数
st = []                     #图中点是否访问过
INF = 0x3f3f3f3f

#有向无环图的拓扑排序  时间复杂度 O(n+m), n 表示点数，m 表示边数
d = []  #每个点的入度
def topsort() -> (bool, int):
    cnt, res = 0, []
    for i in range(1, N + 1):
        if d[i] == 0:
            dq.append(i)
            cnt += 1
    while dq:
        t = dq.pop()
        res.append(t)
        for ne in dfd[t]:
            d[ne] -= 1
            if d[ne] == 0:
                cnt += 1
                dq.append(ne)
    return cnt == N, res    #是否存在拓扑序列，拓扑序列结果

