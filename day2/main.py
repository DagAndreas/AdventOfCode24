input = 'input.txt'

def is_valid_list(lst: list, prev: int, asc = None):
    if not lst:
        return True
    
    if asc == None:
        asc = True if prev < lst[0] else False

    
    if asc:
        if lst[0]-prev < 4 and lst[0]-prev > 0:
            return is_valid_list(lst[1:], asc=True, prev=lst[0])
        
    elif prev - lst[0] < 4 and prev - lst[0] > 0 : 
        return is_valid_list(lst[1:], asc=False, prev=lst[0])    

    return False


def with_dampener(lst: list):
    # exhaustive search through lists and check if they are valid when each one is removed
    safe = False
    for i in range(len(lst)):
        copied_list: list = list.copy(lst)
        copied_list.pop(i)
        safe = safe or is_valid_list(lst=copied_list[1:], prev=copied_list[0], asc=True) or is_valid_list(lst=copied_list[1:], prev=copied_list[0], asc=False)

    return safe
    
safe_star_1 = 0
safe_star_2 = 0
with open(input) as fil:
    for line in fil:
        print(line)
        vals: list = line.replace('\n', '').split(' ')
        vals_int = list(map(int, vals))
        if is_valid_list(vals_int[1:], vals_int[0], None):
            safe_star_1 += 1
        print(f'{line=},{is_valid_list(vals_int[1:], vals_int[0], None)}')
        if with_dampener(vals_int):
            safe_star_2 += 1

print(f'{safe_star_1=}')
print(f'{safe_star_2=}')

# star 1: 220
# star 2: 296