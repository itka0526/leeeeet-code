# m, n = 3, 1

# # split = 2, hub = 5

# c2, c5 = 1, 10

# class Solution:
#     min_cost = float('inf')

#     def rec(self, filled_ports, target, cost):
#         if filled_ports == target:
#             self.min_cost = min(self.min_cost, cost)
#             return
#         if filled_ports > target or cost >= self.min_cost:
#             return
#         # Try using a 2
#         self.rec(filled_ports + 2, target, cost + c2)
#         # Try using a 5
#         self.rec(filled_ports + 5, target, cost + c5)


# n = Solution()

# n.rec(n, m, 0)

# m, n = 3, 1
# c2, c5 = 1, 10
n = int(input())
m = int(input())
c2 = int(input())
c5 = int(input())

def calc(n, m, c2, c5):
    min_cost = [float('inf')]

    def rec(fp, cost):
        if fp >= m:
            min_cost[0] = min(min_cost[0], cost)
            return
        # Try using a 2
        if fp + 2 <= m:
            rec(fp + 2, cost + c2)
        # Try using a 5
        if fp + 5 <= m:
            rec(fp + 5, cost + c5)

    rec(n, 0)
    return min_cost

min_cost = calc(n, m, c2, c5)
print(min_cost)
