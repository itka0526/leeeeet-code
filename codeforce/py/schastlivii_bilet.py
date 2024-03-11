#!/usr/bin/env python3

import sys

r = sys.stdin.readline().strip()
fh = sum(map(int, r[:len(r)//2]))
lh = sum(map(int, r[len(r)//2:]))
print("YES" if fh - lh == 0 else "NO")
