import heapq
from collections import defaultdict, deque


dfd = defaultdict(list)     #邻接表
dq = deque()                #双端队列
N, M = 100, 100             #图中点和边数
st = []                     #图中点是否访问过
INF = 0x3f3f3f3f
dist = []
#时间复杂度 平均情况下 O(m)，最坏情况下 O(nm), n 表示点数，m 表示边数
#st数组表示结点是否在队列
#dq存储有更新的结点（已在队列无需重复入队）
def spfa() -> int:
    dist[1] = 0
    dq.append(1)
    st[1] = True
    while dq:
        t = dq.popleft()
        st[t] = False

        for ne, w in dfd[t]:
            if dist[ne] > dist[t] + w:
                dist[ne] = dist[t] + w
                if not st[ne]:
                    dq.append(ne)
                    st[ne] = True

    return dist[-1]