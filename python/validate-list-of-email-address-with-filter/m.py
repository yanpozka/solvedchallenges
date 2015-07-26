# coding: utf-8

import re
pat = re.compile(r"^[\w\-]+\@[a-z0-9]+(\.[a-z0-9]+)*\.[a-z]{1,3}$")


def valid_email(s):
    st = str(s).strip().lower()
    return (pat.match(st) != None)

N = int(raw_input())

lws = []

for i in xrange(0, N):
    lws.append(raw_input())

l = list(filter(valid_email, lws))

print(sorted(l))
