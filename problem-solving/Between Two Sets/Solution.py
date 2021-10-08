#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'getTotalX' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER_ARRAY a
#  2. INTEGER_ARRAY b
#
from functools import reduce

def getTotalX(a, b):
    def gcd(a,b):
        if b==0:
            return a
        return gcd(b,a%b)
        
    def lcm(a,b):
        return a*b//gcd(a,b)
        
    f1 = reduce(lcm,a)
    f2 = reduce(gcd,b)
    
    count =0
    sus =0
    while sus<=f2:
        sus+= f1
        if f2%sus ==0:
            count+=1
    
    return count
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    m = int(first_multiple_input[1])

    arr = list(map(int, input().rstrip().split()))

    brr = list(map(int, input().rstrip().split()))

    total = getTotalX(arr, brr)

    fptr.write(str(total) + '\n')

    fptr.close()
