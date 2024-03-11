#!/usr/bin/env python3

import sys

next(sys.stdin)
a = set(map(int, sys.stdin.readline().strip().split()))
b = set(map(int, sys.stdin.readline().strip().split()))
x = list(a.intersection(b))
x.sort()
for i in x:
    sys.stdout.write(str(i) + " ")
