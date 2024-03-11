#!/usr/bin/env python3
import sys


a = list(map(int, sys.stdin.readline().strip().split()))
a.sort()
print(a[len(a) - 1] - a[0])
