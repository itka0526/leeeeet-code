#!/usr/bin/env python3

import sys


n1 = int(sys.stdin.readline().strip())
n2 = int(sys.stdin.readline().strip())
if n1 < n2:
    sys.stdout.write("<")
elif n1 > n2:
    sys.stdout.write(">")
else:
    sys.stdout.write("=")
