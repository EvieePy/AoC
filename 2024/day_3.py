import utils

import re


helper = utils.InputHelper(day=3, year=2024)
helper.fetch()


LINES: str = helper.data
# LINES: str = open("EXAMPLE.txt").read().strip()


muls = re.findall(r"mul\((\d+),(\d+)\)|(don't)|(do)", LINES)

cur = True
p1 = 0
p2 = 0

for mul in muls:
    cur = False if mul[2] else True if mul[3] else cur
    answer = int(mul[0]) * int(mul[1]) if mul[0] else 0

    p1 += answer
    if cur:
        p2 += answer


print(f"DAY THREE TOTAL_ONE: {p1}")
print(f"DAY THREE TOTAL_TWO: {p2}")
