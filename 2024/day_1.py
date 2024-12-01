import utils
import collections as ct


helper = utils.InputHelper(day=1, year=2024)
helper.fetch()


LINES: list[str] = helper.as_lines()
TOTAL_ONE: int = 0
TOTAL_TWO: int = 0


left: list[int] = []
right: list[int] = []


for line in LINES:
    splat: list[str] = line.split()
    left.append(int(splat[0]))
    right.append(int(splat[1]))


left.sort()
right.sort()


for index, number in enumerate(left):
    TOTAL_ONE += abs(number - right[index])


print(TOTAL_ONE)


# PART 2

c2 = ct.Counter(right)
for number in left:
    TOTAL_TWO += number * c2.get(number, 0)


print(TOTAL_TWO)
