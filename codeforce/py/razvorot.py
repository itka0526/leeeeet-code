#!/usr/bin/env python3

import sys
n = int(sys.stdin.readline().strip())
x = list(map(int, sys.stdin.readline().strip().split()))
x.reverse()
for i in x:
    sys.stdout.write(str(i) + " ")
