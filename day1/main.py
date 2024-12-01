left_list = []
right_list = []

with open('input.txt') as fil:
    for line in fil:
        nums = line.split()
        left_list.append(int(nums[0]))
        right_list.append(int(nums[1]))

print(left_list)
def value_of_min(lst: list) -> int:
    import math
    min_value = math.inf 

    for num in lst:
        if  num < min_value:
            min_value = num

    if min_value == math.inf:
        SystemExit('Failed in index of min')

    return min_value 

distance = 0

for i in range(len(left_list)):
    min_left_value = value_of_min(left_list)
    left_list.remove(min_left_value)

    min_right_value = value_of_min(right_list)
    right_list.remove(min_right_value)

    distance_between: int = min_right_value - min_left_value
    if distance_between < 0:
        distance_between = distance_between * -1

    distance += distance_between

print(f'{distance=}')