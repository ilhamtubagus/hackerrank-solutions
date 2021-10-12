#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'timeConversion' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#

def timeConversion(s):
    # Write your code here
    t = s.split(":")
    if(s[-2:]== 'PM'):
        if t[0] != "12":
            t[0] = str(int(t[0])+12)
            
    else:
        if t[0] == "12":
            t[0] = "00"
            
    t24 = ':'.join(t)
    return str(t24[:-2])


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = timeConversion(s)

    fptr.write(result + '\n')

    fptr.close()
