#!/usr/bin/env python3
import sys

n, m, q = map(int, sys.stdin.readline().strip().split())
col_name = sys.stdin.readline().strip().split()

table = []
for _ in range(n):
    row_val = list(map(int, sys.stdin.readline().strip().split()))
    table.append({col_name[i]: row_val[i] for i in range(m)})

for _ in range(q):
    col_name, operator, value = sys.stdin.readline().strip().split()
    value = int(value)
    if operator == '>':
        table = [row for row in table if row[col_name] > value]
    elif operator == '<':
        table = [row for row in table if row[col_name] < value]

total_sum = sum(sum(row.values()) for row in table)
sys.stdout.write(str(total_sum))
