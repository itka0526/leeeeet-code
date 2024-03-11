#!/usr/bin/env python3

import sys

a, b, c = map(int, sys.stdin.readline().strip().split())
sys.stdout.write("YES" if a * b == c else "NO")
