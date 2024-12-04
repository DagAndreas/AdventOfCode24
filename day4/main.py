import pandas as pd
import numpy as np

testdata = []
df = pd.read_csv('input.txt', header=None, names=['data'])


df = pd.DataFrame(df['data'].apply(list).tolist())

df.columns = [chr(97 + i) for i in range(df.shape[1])] 

def find_how_many_matches(string: str, pattern: str = 'XMAS'):
    print(f'found {string.count(pattern)} in {string}')
    return string.count(pattern) 

a = 'abcxmasabcxmasabs'

ans = find_how_many_matches(a, 'xmas')
print(f'{ans=}')

def process_np_array(arr) -> int:
    # values are of type array and can be used as lists when joining. 
    line_string = ''.join(arr)
    
    # nparrays can be flipped by:
    # arr[::-1]
    reverse_string = ''.join(arr[::-1])

    return find_how_many_matches(line_string) + find_how_many_matches(reverse_string)

def process_row(row) -> int:
    # row[0] is row index
    # row[1] are values
    line = row[1].values
    return process_np_array(line)

sum = 0
# read all horisontally
for row in df.iterrows():
    sum += process_row(row)

# read all vertically
for row in df.T.iterrows():
    sum += process_row(row)


# read desc diagonals
hor_length, vertLength = df.shape
flipped_df = np.fliplr(df.values)

# do all diagonals on the horizontal axis 
for i in range(1, hor_length):
    diag_array = np.diag(df.values, k=i)
    sum += process_np_array(diag_array)

    diag_array = np.diag(flipped_df, k=i)
    sum += process_np_array(diag_array)

# the center diagonal
sum += process_np_array(np.diag(df.values, k=0))
sum += process_np_array(np.diag(flipped_df, k=0))

# desc diagonals
for i in range(1, vertLength):
    diag_array = np.diag(df.values, k=-i)
    sum += process_np_array(diag_array)

    diag_array = np.diag(flipped_df, k=-i)
    sum += process_np_array(diag_array)


# read asc diagonals


print(f'{sum=}')

# star1: 2685
