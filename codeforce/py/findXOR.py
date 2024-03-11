#!/usr/bin/env python3
import sys


def einput(): return sys.stdin.readline().strip()


n = int(einput())
a = [0] * n
k = 0
for e in einput().split(" "):
    a[k] = int(e)
    k += 1
b = [0] * (n + 1)
q = int(einput().strip())
for i in range(0, n):
    b[i + 1] = b[i] ^ a[i]

for i in range(0, q):
    n1, n2 = map(int, einput().split(" "))
    if n1 - 1 > 0:
        print(b[n2] ^ b[n1 - 1])
    else:
        print(int(b[n2]))
