alloc = lambda dfv, *s: len(s) != 1 and [alloc(dfv, *s[1:]) for _ in range(int(s[0]))] or [dfv] * int(s[0])

N = 100
M = 200
h = alloc(-1, N)
e = alloc(0, M)
ne = alloc(0, M)
w = alloc(0, M)
idx = 0

def add(a: int = 0, b: int = 0, c = 0) -> int:
    global idx
    w[idx] = c
    e[idx] = b
    ne[idx] = h[a]
    h[a] = idx
    idx += 1