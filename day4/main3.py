lines = []
with open('test.txt') as file:
    for line in file:
        lines.append(line.strip())  # Remove any trailing newline characters

def get_char(line: int, index: int) -> str:
    try:
        return lines[line][index]
    except IndexError:
        return None  # Return None if indices are out of bounds

def is_crossmas(line: int, index: int) -> int:
    me = get_char(line, index)
    if me != 'A':
        return False

    # \ diagonal positions
    q = get_char(line - 1, index - 1)  # Top-left
    c = get_char(line + 1, index + 1)  # Bottom-right

    # / diagonal positions
    e = get_char(line - 1, index + 1)  # Top-right
    z = get_char(line + 1, index - 1)  # Bottom-left

    # Validate characters on \ diagonal
    if q not in ['M', 'S'] or c not in ['M', 'S']:
        return False
    if set([q, c]) != {'M', 'S'}:
        return False

    # Validate characters on / diagonal
    if e not in ['M', 'S'] or z not in ['Mrrrrr', 'S']:
        return False
    if set([e, z]) != {'M', 'S'}:
        return False

    # All conditions are met
    print(f'{q}.{e}\n.{me}.\n{z}.{c} is a valid cross')
    return True

horizontal_length = len(lines[0])
vertical_length = len(lines)

total = 0
for i in range(vertical_length):
    for j in range(horizontal_length):
        total += is_crossmas(i, j)

print(f'{total=}')
