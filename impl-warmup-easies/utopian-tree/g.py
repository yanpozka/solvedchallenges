#!/usr/bin/python

r = []
r.append(1)

for ix in xrange(1, 61):
    if ix % 2 == 0:
        r.append (r[ix-1] + 1)
    else:
        r.append(r[ix-1] * 2)

print r
