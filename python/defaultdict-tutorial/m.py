#coding: utf-8

from collections import defaultdict

dic = defaultdict(list)

parts = raw_input().split(" ")

N, M = int(parts[0]), int(parts[1])

for ix in range(1, N+1):
    a = raw_input()
    dic[a].append(ix)

for ix in range(0, M):
    b = raw_input()
    ok = False
    if len(dic[b]) > 0:
        ls = sorted(dic[b])
        print ' '.join(str(e) for e in ls)
    else:
        print -1
