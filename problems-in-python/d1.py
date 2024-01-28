class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# 3 - общее количество
# 0 - тут всегда ноль, потому что нулевой элемент
# 0 - это значит, что у первого элемента, родитель 0
# 1 - у второго элемента родитель 1

# 1. `8` - This is the total number of repositories (`N`), which means there are 8 repositories in total.

# 2. `0` - This line represents the first repository (index 0), and it indicates that the first repository (index 0) is the root or initial repository. It has no parent and serves as the starting point.

# 3. `1` - This line represents the second repository (index 1), and it indicates that the second repository (index 1) inherits changes from the repository with index 1. In other words, the second repository is a child of the first repository (index 1).

# 4. `2` - This line represents the third repository (index 2), and it indicates that the third repository (index 2) inherits changes from the repository with index 2. In other words, the third repository is a child of the second repository (index 1).

# 5. `0` - This line represents the fourth repository (index 3), and it indicates that the fourth repository (index 3) is a direct child of the root repository (index 0). In other words, it inherits changes from the root repository.

# 6. `4` - This line represents the fifth repository (index 4), and it indicates that the fifth repository (index 4) inherits changes from the repository with index 4. In other words, the fifth repository is a child of the repository with index 4.

# 7. `5` - This line represents the sixth repository (index 5), and it indicates that the sixth repository (index 5) inherits changes from the repository with index 5. In other words, the sixth repository is a child of the repository with index 5.

# 8. `6` - This line represents the seventh repository (index 6), and it indicates that the seventh repository (index 6) inherits changes from the repository with index 6. In other words, the seventh repository is a child of the repository with index 6.

# 9. `4` - This line represen}{]]ts the eighth repository (index 7), and it indicates that the eighth repository (index 7) inherits changes from the repository with index 4. In other words, the eighth repository is a child of the repository with index 4.

# n = 8
# nodes = [0, 1, 2, 0, 4, 5, 6, 4]
# N = int(input())
# inheritance = [int(input()) for _ in range(N)]

# # Initialize counts for each repository
# counts = [0] * N

# # Iterate through inheritance descriptions and update counts
# for i in range(N):
#     parent = inheritance[i]
#     counts[parent] += 1  # Increment count for the parent repository

# # Find the repository with the maximum count
# max_count = max(counts)
# repository_with_max_count = counts.index(max_count)

# # Print the output (1-based index)
# print(repository_with_max_count + 1)

# k, n, m = map(int, input().split()) # количество тротуаров, количество раскопок и количество укладок плитки
 
k, n, m = 5, 4, 3 # m - budget
tiles_needed = [0] * k
unhappy = 0



# for _ in range(n):
#     d, w = map(int, input().split())
#     tiles_needed.append((d, w))
