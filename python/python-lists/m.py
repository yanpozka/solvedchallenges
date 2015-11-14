# coding: utf-8

N = int(raw_input())
mialist = []

while N > 0:
    N -= 1
    parts = raw_input().split()
    # print parts
    
    l = len(parts)

    if l == 1:
        if parts[0] == "print":
            print mialist
        else:
            getattr(mialist, parts[0])()
    elif l == 2:
        fn = getattr(mialist, parts[0])
        val = parts[1]
        try:
            fn(int(parts[1]))
        except TypeError as te:
            fn(parts[1])
    else:
        fn = getattr(mialist, parts[0])
        ix = int(parts[1])
        try:
            fn(ix, int(parts[2]))
        except TypeError as te:
            fn(ix, parts[2])
 
