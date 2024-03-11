#!/usr/bin/env python3

import sys


n = int(sys.stdin.readline().strip())
a, b = 0, 0
for i in range(0, n):
    if int(sys.stdin.readline().strip()):
        a += 1
    else:
        b += 1

print(min(a, b))
