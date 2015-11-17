import re

regExp = r'\#[0-9A-Fa-f]{3,6}'

N = int(raw_input())

source = ""

while N > 0:
    N -= 1
    source += raw_input() 
    source += " "

parts = source.split(";")

for p in parts:
    ix = p.find("{")
    if ix != -1:
        p = p[ix+1:]
    
    colors = re.findall(regExp, p)
    for c in colors:
        print c


