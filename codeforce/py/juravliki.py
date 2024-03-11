#!/usr/bin/env python3
import sys

# n1, n2, n3 = n1 - n3 == 0 && n1 + n2 + n3 = n && n2 = n1 * 2
# n2 = n - n1 - n3 &&  n2 = n1 * 2
# n3 = 2 * (n1 + n1)
# n1 + n2 + n3 -> n1 + n2 + 2*n1 + 2*n2 = 3*n1 + 3*n2; n1 = n2 -> 6*x
n = int(sys.stdin.readline().strip())
print(n // 6,  4 * (n // 6), n // 6)
