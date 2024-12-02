import utils


helper = utils.InputHelper(day=2, year=2024)
helper.fetch()


LINES: list[str] = helper.as_lines()
# LINES: list[str] = open("EXAMPLE.txt").read().strip().splitlines()


def check(nums):
    prev = None
    prev_dist = 0

    for num in nums:
        n = int(num)

        if not prev:
            prev = n
            continue

        dist = prev - n
        if abs(dist) > 3:
            return False

        if dist == 0:
            return False
        elif dist > 0 and prev_dist < 0:
            return False
        elif dist < 0 and prev_dist > 0:
            return False

        prev_dist = dist
        prev = n

    return True


P1 = 0
for line in LINES:
    nums = line.split()
    P1 += check(nums)


P2 = 0
for line in LINES:
    nums = line.split()
    P2 += check(nums) or any(check(nums[:n] + nums[n + 1 :]) for n in range(len(nums)))


print(P1)
print(P2)
