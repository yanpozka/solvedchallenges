# coding: utf-8

N = int(raw_input())

words = {}
lws = []

for i in xrange(0, N):
    s = raw_input()
    if s not in words:
        words[s] = 0
        lws.append(s)
    words[s] += 1

print len(lws)

print " ".join([str(words[k]) for k in lws])
