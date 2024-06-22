def calculate_max_tips(schedules: list):
    schedules = sorted(schedules, key=lambda x: (x[1], x[0]))
    current_time = 0
    total_tips = 0

    for a, b in schedules:
        delivery_time = current_time + b
        total_tips += a - delivery_time
        current_time = delivery_time

    return total_tips


def update_schedule_and_recalculate(schedules, update_index, new_a, new_b):
    schedules[update_index] = (new_a, new_b)
    return calculate_max_tips(schedules)


n, q = map(int, input().split())

schedules = []
for _ in range(n):
    ab = tuple(map(int, input().split()))
    schedules.append(ab)

result = [calculate_max_tips(schedules)]

for _ in range(q):
    i, a, b = map(int, input().split())
    result.append(update_schedule_and_recalculate(schedules, i - 1, a, b))

print("\n".join(map(str, result)))
