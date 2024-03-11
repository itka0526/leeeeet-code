#!/usr/bin/env python3

import sys

n = int(sys.stdin.readline().strip())
k = 0
arr = [0] * n
while (k < n):
    next(sys.stdin)
    arr[k] = list(map(int, sys.stdin.readline().strip().split()))
    k += 1

total = 0

for i in range(len(arr)):
    for j in range(i + 1, len(arr)):
        prefix = []
        for a, b in zip(arr[i], arr[j]):
            if a == b:
                prefix.append(a)
            else:
                break
        total += len(prefix)

sys.stdout.write(str(total))
