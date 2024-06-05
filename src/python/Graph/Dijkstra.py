from collections import defaultdict
from heapq import heappop, heappush
#有向无环正权图,朴素版  时间复杂是 O(n^2+m), n 表示点数，m 表示边数
#dist[N]  // 存储1号点到每个点的最短距离,初始化为 INF
#求1号点到n号点的最短路，如果不存在则返回-1
N = 100
st = []
dist = []
dfd = defaultdict(list)
INF = float('inf')

def dijkstra() -> int:
    dist[1] = 0
    for i in range(N - 1):
        t = -1     # 在还未确定最短路的点中，寻找距离最小的点
        for j in range(1, N + 1):
            if not st[j] and (t == -1 or dist[t] > dist[j]):
                t = j
        # 用t更新其他点的距离
        for (ne, w) in dfd[t]:
            dist[ne] = min(dist[ne], dist[t] + w)
        st[t] = True
    return -1 if dist[N] == INF else dist[N]


#堆优化版  时间复杂度 O(mlogn), n 表示点数，m 表示边数
heap = []
def dijkstra() -> int:
    dist[1] = 0
    heappush(heap, (0, 1))  # 最短距离, 节点编号
    while heap:
        distance, ver = heappop(heap)
        if st[ver]: continue
        st[ver] = True
        for ne, w in dfd[ver]:
            if dist[ne] > distance + w:
                dist[ne] = distance + w
                heappush(heap, (dist[ne], ne))
    return -1 if dist[N] == INF else dist[N]