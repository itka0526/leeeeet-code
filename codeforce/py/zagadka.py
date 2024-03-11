#!/usr/bin/env python3

import sys

# a + b = x && a * b = y
x, y = map(int, sys.stdin.readline().strip().split())

try:
    for i in range(0, 1001):
        for j in range(0, 1001):
            if i + j == x and i * j == y:
                print(i, j)
                raise StopIteration
except StopIteration:
    pass
