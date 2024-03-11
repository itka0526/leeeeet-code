#!/usr/bin/env python3

import sys

l = sys.stdin.readline().strip()
m = 0
lc = 0
for c in l:
    if c == "0":
        lc += 1
    else:
        lc = 0
    if lc > m:
        m = lc

print(m)
