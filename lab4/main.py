from queue import Queue


graph = [(1, 0), (1, 3), (6, -1), (6, -6), (-3, 0), (-3, 1)]

n = len(graph)


def m(v, w): return abs(v[0]-w[0])+abs(v[1]-w[1])


edges = {}
for i in range(n):
    for j in range(n):
        v, w = graph[i], graph[j]
        c = m(v, w)
        if i == j or (j, i) in edges:
            continue
        edges[i, j] = c

Q = Queue()
for e in sorted(edges, key=edges.get):
    Q.put(e)

name = [0]*n
next = [0]*n
size = [0]*n


def merge(v, w, p, q):
    name[w] = p
    u = next[w]
    while name[u] != p:
        name[u] = p
        u = next[u]
    size[p] = size[p] + size[q]
    next[v], next[w] = next[w], next[v]


for v in range(n):
    name[v], next[v], size[v] = v, v, 1

tree = []
while len(tree) != n - 1:
    (v, w) = Q.get()
    p, q = name[v], name[w]
    if p != q:
        if size[p] > size[q]:
            merge(w, v, q, p)
        else:
            merge(v, w, p, q)
        tree.append((v, w))

print(tree)

print(sum([m(graph[e[0]], graph[e[1]]) for e in tree]))
