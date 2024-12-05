rules = {}
found_nums = []

def is_in_rule(x, y):
    try:
        return rules[f'{x}|{y}']
    except KeyError:
        return False

middle_num_sum = 0
incorrect_middle_sum = 0
retry_sum = 0
with open('input.txt') as file:
    for line in file:
        if line == '\n':
            continue

        # ordering rules
        if line.count('|') > 0:

            rules[f'{line.strip()}'] = True


            continue

        # check the ordering
        is_in_order = True
        split_line = line.strip().split(',')
        # iter through rules and check if they are valid.
        for i in range(len(split_line) - 1):
            for j in range(i+1, len(split_line)):
                is_in_order = is_in_order and is_in_rule(split_line[i], split_line[j])
        if is_in_order:
            middle_num_sum += int(split_line[int(len(split_line)/2)])
            continue

        # part 2
        if not is_in_order:
           for i in range(0, len(split_line) - 1):
                for j in range(i+1, len(split_line)):
                    x = split_line[i]
                    y = split_line[j]
                    if not is_in_rule(x, y):
                        print(f'{x=}, {y=} it not in order')         
                        rem = split_line.pop(j)
                        split_line.insert(i, rem)                

        incorrect_middle_sum += int(split_line[int(len(split_line)/2)])

print(f'{middle_num_sum=}')
print(f'{incorrect_middle_sum=}')
print(f'{retry_sum=}')
# star1: 5391

# star2: ans = 6142 > 6110 > 5786