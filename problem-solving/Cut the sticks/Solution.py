#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'cutTheSticks' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts INTEGER_ARRAY arr as parameter.
#

def cutTheSticks(arr):
    # Write your code here
    ptr = arr.index(max(arr))
    sticks = []
    while arr[ptr]!=0:
        count = 0
        y = []
        cut = findmin(arr)
        for i in range(len(arr)):
            if arr[i] >=cut:
                count+=1
                arr[i] = arr[i]-cut
        sticks.append(count)
        
    return sticks

def findmin(x):
    mins = 1000
    for j in x:
        
        if j < mins and j != 0:
            mins = j
    return mins
                

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    result = cutTheSticks(arr)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
