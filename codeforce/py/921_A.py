#!/usr/bin/env python3
import sys

t = int(sys.stdin.readline().strip())
a = "abcdefghijklmnopqrstuvwxyz"

while t > 0:
    n, k, m = map(int, sys.stdin.readline().strip().split(" "))
    s1 = sys.stdin.readline().strip()

    s2 = a[:k]
    count = [0] * 26
    required = m
    for c in s1:
        count[c - int('a')] += 1
    l, r = 0, 0
    while (r < k):
        if (count[s2[r] - int('a')] >= 0):
            count -= 1
            required -= 1
        while (required == 0):
            if r - l + 1 == len(s1):
                print("Yes")
            if (count[s2[l + 1] - int('a')] > 0):
                l += 1
                required += 1
                count += 1
        r += 1
    print("No")
    t -= 1
