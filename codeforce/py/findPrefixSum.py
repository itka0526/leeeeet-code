#!/usr/bin/env python3
import sys

n = int(sys.stdin.readline().strip())
b = [0] * (n + 1)
i = 0
for e in sys.stdin.readline().strip().split():
    b[i + 1] = int(e) + b[i]
    i += 1
q = int(sys.stdin.readline().strip())
for _ in range(q):
    n1, n2 = map(int, sys.stdin.readline().strip().split(" "))
    if n1 - 1 < 0:
        sys.stdout.write(str(b[n2] - b[n1]) + "\n")
    else:
        sys.stdout.write(str(b[n2] - b[n1 - 1]) + "\n")
