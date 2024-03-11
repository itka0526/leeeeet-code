#!/usr/bin/env python3

import sys
from collections import Counter

n, m, q = map(int, sys.stdin.readline().strip().split())
a = Counter(map(int, sys.stdin.readline().strip().split()))
b = Counter(map(int, sys.stdin.readline().strip().split()))

for _ in range(q):
    [type, player, card] = sys.stdin.readline().strip().split()
    type = int(type)
    card = int(card)
    if player == 'A':
        a[card] += type
    else:
        b[card] += type

    cards = set(a.keys()).union(b.keys())
    c = 0
    for card in cards:
        c += abs(a[card] - b[card])
    sys.stdout.write(str(c) + " ")

# 2 1 0 1 2 3 2 1 0 1
