#!/usr/bin/env python3
from datetime import datetime
import sys
import calendar

# (1≤year1≤9999,1≤month1≤12,1≤day1≤31,0≤hour1≤23,0≤min1≤59,0≤sec1≤59)
start_date = list(map(int, sys.stdin.readline().strip().split()))
end_date = list(map(int, sys.stdin.readline().strip().split()))


start_datetime = datetime(
    start_date[0], start_date[1], start_date[2], start_date[3], start_date[4], start_date[5])
end_datetime = datetime(
    end_date[0], end_date[1], end_date[2], end_date[3], end_date[4], end_date[5])


delta = end_datetime - start_datetime
delta_days = delta.days

leap_days = 0

for year in range(start_date[0], end_date[0] + 1):
    if calendar.isleap(year) and start_datetime < datetime(year, 3, 1):
        leap_days += 1

delta_days -= leap_days

print(delta_days, delta.seconds)
# 17 96
# 2923033 79555
