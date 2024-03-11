#!/usr/bin/env python3

import sys

n = int(sys.stdin.readline().strip())

a = [0] * n
i = 0
for e in sys.stdin.readline().strip().split():
    a[i] = int(e)
    i += 1
print(min(a), max(a))
