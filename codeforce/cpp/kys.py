file = open("./shit", "+w")

s = f"""
a = int(input())
b = int(input())

if (a == 1 and b == 1):
    print(2)
"""

s1 = [
    f"elif (a == {i} and b == {i + 1}):\n\tprint({i + i + 1})\n" for i in range(100)]

s1 = "".join(s1)

s2 = f"else:\n\t print(0)"

s3 = s + s1 + s2

file.write(s3)
