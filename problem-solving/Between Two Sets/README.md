If you confused after reading this problem for the first time, dont worry i will explain what is this problem about.

You are given 2 arrays of integer, You need to find integers which met this condition:
- Multiple of first array
- Factors of second array

At test case#1 you are given:
a = [2, 6]
b = [24, 36]

First thing you want to do is to to get multiple of first array:

2 = 2, 4, **6**, 8, 10, **12**, ...
6 = **6**, **12**, 18, 24, 30, ...

Then you need to find factors of second array:

24 = 1, 2, 3, 4, **6**, 8, **12**, 24

36 = 1, 2, 3, 4, **6**, 9, **12**, 18, 36

Final step is to count how many numbers that are the appear at every element of first multiple of the first array and factors of the second array, which is 2 numbers (6 and 12)

lets take a look at test case#2.
a = [2, 4]
b = [16, 32, 96]

Multiple of first array is:
2 = 2, **4**, 6, **8**, 10, 12, 14, **16**, ...
4 = **4**, **8**, 12, **16**, 20, ...

Factors of second aray is:
16 = 1, 2, **4**, **8**, **16**

32 = 1, 2, 3, **4**, 6, **8**, **16**
     
96 = 1, 2, 3, **4**, 6, **8**, 12, **16**, 24, 32, 48, 96
     
so the output of test case#2 is 3, because number 4, 8 and 16.
