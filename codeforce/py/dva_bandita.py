#!/usr/bin/env python3

# 4 + 7 = 11 - 1; -1 because both of them shot the last one
# 10 / 2 = 5; Ideally each one should have shot at 5 bottles.
# 5 + (5 - 4) = 6
# 5 + (5 - 7) = 3

import sys
larry, garry = map(int, sys.stdin.readline().strip().split(" "))
total = larry + garry - 1
ideal = total / 2
print(int(ideal + (ideal - larry)), int(ideal + (ideal - garry)))
