from collections import defaultdict, deque
"""
深搜打表，跳跃查询
O((n+m)logn)
"""
# 最近公共祖先(Lowest Common Ancestor, LCA)
# 倍增算法 求LCA
# dep[u] 存u结点的深度
# fa[u][i] 从u向上跳2^i层的祖先结点  i=0,1,2...

# 1) dfs初始化 dep/fa数组
# 2) 1.求uv的LCA，先将 uv 跳到同一层，两点深度查差为y，将更深的点以y的2进制方式跳
#    2.将uv一起跳到LCA的下一层

N = 10 ** 5
dep, fa = [N], [N][20]
dfd = defaultdict(list)

#结点下标从 1 开始，数以父结点数组形式给出(下标从0开始)
def init() -> None:
    n, m = 10 * 5, 20 - 1
    arr = [] # i 下标的父结点为 arr[i]
    # fa 的第 0 行为哨兵结点
    fa = [[0] * (m + 1)] + [[p + 1] + [0] * m for p in arr]
    for i in range(m):
        for x in range(n):
            fa[x][i + 1] = fa[fa[x][i]][i]
"""
深搜初始化：dep1全为0
宽搜初始化：dep[0]=0,dep[root]=1,其余为正无穷
"""
#结点编号从 1 开始
def dfs(u: int = 0, p: int = 0) -> None:
    dep[u] = dep[p] + 1
    #向上跳2的 1,2,3 次方步的祖先结点
    fa[u][0] = p
    for i in range(1, 20):
        fa[u][i] = fa[fa[u][i - 1]][i - 1]

    for ne in dfd[u]:
        if ne != p: dfs(ne, u)

# bfs 初始化
# 注： 以最短路的思想，dep数组初始化为正无穷，表示未访问过
def bfs(root: int = 0):
    dq = deque()
    dq.append(root)
    dep[0] = 0      #0结点(哨兵)深度为0
    dep[root] = 1
    while dq:
        t = dq.popleft()
        for ne in dfd[t]:
            if dep[ne] > dep[t] + 1:
                dep[ne] = dep[t] + 1
                dq.append(ne)
                fa[ne][0] = t
                for j in range(1, 20):
                    fa[ne][j] = fa[fa[ne][j - 1]][j - 1]

# 关键  1) 为了保证跳到同一层，一定要比较深度
#       2) 为了保证跳到最近公共祖先的下一层，一定要比较他们的公共祖先
def lca(u: int = 0, v: int = 0) -> int:
    #让 u 跳
    if dep[u] < dep[v]: u, v = v, u

    # 先跳到同一层
    for i in range(19, -1, -1):
        if dep[fa[u][i]] >= dep[v]:
            u = fa[u][i]
    # u v 在一条链上
    if u == v: return u

    # u v 一起跳到LCA的下一层
    for i in range(19, -1, -1):
        if fa[u][i] != fa[v][i]:
            u, v = fa[u][i], fa[v][i]
    # u v 停在最近公共祖先的下一层，返回再跳1步的值
    return fa[u][0]

# fa[u][i] != fa[v][i] 保证了u v 不会跳到公共祖先上，因为有可能跳到非最近的公共祖先
