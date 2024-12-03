import enum 
test = 'xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]'
test = 'xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))'

legal_characters = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0, '(', ')', ',']
legal_nums = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

do = True
def is_num(c: str):
    try:
        return int(c) in legal_nums
    except TypeError:
        return False

def is_valid(c: str) -> bool:
    return c in legal_characters

def parse_num(line: str, index: int) -> tuple[int, int, str]:
    # can max be 3 numbers
    num = 0
    for i in range(4):
        line_char = line[index+i]
        # found comma. Return
        if line_char == ',' or line_char == ')':
            if (num == 0):
                raise ValueError
            
           # num is not zero and comma has been found. We can return it. 
            return (num, index + i, line_char)

        if not is_num(line_char):
            raise ValueError
        if line_char == '2' or line_char == '1':
            print('hei')
        num = (num*10 + int(line_char))        



def parse_pran(line: str, index: int) -> int:

    num1 = None
    num2 = None
    # parse number and get next index
    num1, comma_index, separator_char = parse_num(line, index)

    # comma
    if separator_char != ',':
        raise ValueError
    
    # parse number and assert it ends with closing param
    num2, comma_index, separator_char = parse_num(line, comma_index+1)

    # check delim
    if separator_char != ')':
        raise ValueError

    print(f'found {num1} x {num2} = {num1*num2}')
    return num1 * num2 


def get_mul_from_string(line: str, do: bool = do) -> int:
    sum = 0

    # iter through the len
    for i in range(len(line)):
        line_char = line[i]
        if len(line) - i < 3:
            continue

        if line_char != 'm' and line_char != 'd':
            continue

        if line_char == 'm':

            if line[i+1] != 'u':
                continue

            if line[i+2] != 'l':
                continue

            if line[i+3] != '(':
                continue

            # parse from opening to closing paranthesis
            try:
                if do:
                    sum += parse_pran(line, index=i+4)
            except ValueError:
                continue
        # Check if enable or disable do
        if line_char == 'd':
                
            if line[i+1] != 'o':
                continue
        
            # do | don't 
            if line[i+2] == '(' and line[i+3] == ')':
                do = True
            
            elif line[i+2] == 'n' and line[i+3] == '\'' and line[i+4] == 't' and line[i+5] == '(' and line[i+6] == ')':
                do = False
            

    return sum
        

# ans = get_mul_from_string(test)
# print(f'{ans=}')

full_line = ''

with open('input.txt') as file:
    for line in file:
        full_line += line


total = 0
total = get_mul_from_string(full_line)

print(f'{total=}')

# star1: 161085926
# star2: 82045421