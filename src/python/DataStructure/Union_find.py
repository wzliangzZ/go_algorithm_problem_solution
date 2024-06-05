

bet = lambda a, b: a <= b and range(a, b + 1) or range(a, b - 1, -1)

"""维护size的并查集：(朴素并查集去掉所有size有关操作)"""
p, size = [], []
n = 100
#p[]存储每个点的祖宗节点, size[]只有祖宗节点的有意义，表示祖宗节点所在集合中的点的数量

def find(x: int) -> int:
    if p[x] != x: p[x] = find(p[x])
    return p[x]


for i in bet(1, n):
    p[i] = i
    size[i] = 1

#合并a和b所在的两个集合，a并入b：
# size[find(b)] += size[find(a)]
# p[find(a)] = find(b)


"""维护到祖宗节点距离的并查集："""
p, d = [], []
#p[]存储每个点的祖宗节点, d[x]存储x到p[x]的距离

#返回x的祖宗节点
def find(x: int) -> int:
    if (p[x] != x):
        u = find(p[x])
        d[x] += d[p[x]]
        p[x] = u
    return p[x]


#初始化，假定节点编号是1~n
for i in range(1, n):
    p[i] = i
    d[i] = 0


#合并a和b所在的两个集合：
# p[find(a)] = find(b)
# d[find(a)] = distance
#根据具体问题，初始化find(a)的偏移量
