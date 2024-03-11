#!/usr/bin/env python3
import sys
next(sys.stdin)
market = list(map(int, sys.stdin.readline().strip().split()))

# max_profit = [0]
# l = len(market)
# def rec(i, p, n):
#     if p > max_profit[0]:
#         max_profit[0] = p
#     if i > l - 1:
#         return
#     # Keep: i++, p  , n++
#     rec(i + 1, p, n + 1)
#     # Sell: i++, p++, n = 1
#     rec(i + 1, p + market[i] * n, 1)
# rec(0, 0, 1)
# print(max_profit[0])

pos = len(market) - 1
c = [0] * len(market)
c[pos] = 1
i = len(market) - 2
# max gets incremented by 1 each time i moves to the left
# if new max gets found, max gets updated to be that i
# in the end we get a list of numbers that resemble the worth buying days
# [ 73 31 96 24 46 ]
# [ 0  0  3  0  2  ]
# 96*3 + 46*2 = 380
# I could not come out with this solution but oh well... My recursion solution is very self-explanatory
# But its time limit exceeds :(

while (i >= 0):
    if (market[pos] >= market[i]):
        c[pos] += 1
    else:
        # Update the pos to the max
        pos = i
        c[pos] = 1
    i -= 1
sum = 0
for i in range(len(c)):
    sum += c[i] * market[i]
print(sum)
