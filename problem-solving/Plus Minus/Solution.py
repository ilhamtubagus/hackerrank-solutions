#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'plusMinus' function below.
#
# The function accepts INTEGER_ARRAY arr as parameter.
#

    
def plusMinus(arr):
    npos = 0
    nneg = 0
    nnol = 0
    
    for i in arr:
        if i > 0:
            npos+=1
        elif i < 0:
            nneg+=1
        else:
            nnol+=1
            
    print(npos/len(arr)) 
    print(nneg/len(arr))
    print(nnol/len(arr))
    
if __name__ == '__main__':
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
