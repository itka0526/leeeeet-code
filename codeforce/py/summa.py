#!/usr/bin/env python3

import sys
n = int(sys.stdin.readline().strip())
if n > 0:
    print(n*(n + 1) // 2)
else:
    print(-n*(n - 1) // 2 + 1)
