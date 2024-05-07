def solve():
    import sys

    data = int(sys.stdin.readline().strip())

    # index = 0
    # t = int(data[index])
    # index += 1
    # results = []
    # print("GOT: ", t)
    # for _ in range(t):
    #     n = int(data[index])
    #     k = int(data[index + 1])
    #     index += 2

    #     if k == 1:
    #         results.append(f"{n-1}")
    #         results.append(" ".join(str(i) for i in range(2, n + 1)))
    #         continue

    #     a = []
    #     current_sum = 0
    #     i = 1
    #     # Генерируем числа до тех пор, пока не превысим n и не обойдем k
    #     while current_sum < n:
    #         if current_sum + i != k and current_sum + i <= n:
    #             a.append(i)
    #             current_sum += i
    #         i += 1

    #     results.append(str(len(a)))
    #     results.append(" ".join(map(str, a)))

    # sys.stdout.write("\n".join(results) + "\n")


solve()
