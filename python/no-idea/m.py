# coding: utf-8

parts = raw_input().split()

N, M = int(parts[0]), int(parts[1])

nums = map(int, raw_input().split()[:N])

A = {}
B = {}

parts = raw_input().split()
for p in parts:
    if p != '':
        A[int(p)] = True

parts = raw_input().split()
for p in parts:
    if p != '':
        B[int(p)] = True

h = 0
for n in nums:
    if n in A:
        h += 1
    else:
        if n in B:
            h -= 1

print h
