data = []
with open('input.txt') as file:
    data = file.read().splitlines()

def get_char(line: int, index: int) -> str:
    try:
        return data[line][index]
    # catch out of bounds
    except:
        return False

def is_crossmas(x: int, y: int) -> int:
    # print(f'found char: {get_char(line, index)}')
    me = get_char(x, y)
    if me != 'A':
        return False

    # \ dia
    q = get_char(x-1, y-1)
    c = get_char(x+1, y+1)
    dia_mas = q == 'S' and c == 'M' or q == 'M' and c == 'S'
    if not dia_mas:
        return False
    
    
    
    # / dia
    z = get_char(x+1, y-1)
    e = get_char(x-1, y+1)
    dia_mas = z == 'S' and e == 'M' or z == 'M' and e == 'S'
    if not dia_mas:
        return False
    
    # if q == 'M' and c == 'S':
        # print(f'1{q}{me}{c}')
    # elif q == 'S' and c == 'M':
        # print(f'2{c}{me}{q}')
    # else:
        # print('panic1!')

    # if z == 'M' and e == 'S':
    #     print(f'3{z}{me}{e}')
    
    # elif z == 'S' and e == 'M':
    #     print(f'4{e}{me}{z}')
    # else:
    #     print('panic2!')

    # # Center is A and both / and \ are valid
    # print(f'{q}.{e}\n.{me}.\n{z}.{c} is a valid cross')
    return True

pass

horizontal_length = len(data[0])
vertical_length = len(data)

total = 0
for i in range(vertical_length-1):
    for j in range(horizontal_length-1):
        # c = get_char(i, j)
        # print(f'{i=}, {j=}, {c=}')
        found = is_crossmas(i, j)
        if found:
            total += 1
print(f'{total=}')




    


