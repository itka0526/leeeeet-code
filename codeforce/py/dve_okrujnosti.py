#!/usr/bin/env python3

import sys
x1, y1, r1 = map(int, sys.stdin.readline().strip().split())
x2, y2, r2 = map(int, sys.stdin.readline().strip().split())

if (r1 - r2) ** 2 <= (x2 - x1) ** 2 + (y2 - y1) ** 2 <= (r1 + r2) ** 2:
    print("YES")
else:
    print("NO")
