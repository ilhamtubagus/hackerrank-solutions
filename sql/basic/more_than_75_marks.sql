// https://www.hackerrank.com/challenges/more-than-75-marks

select name from students where marks > 75 order by RIGHT(name, 3) ASC, ID ASC;
