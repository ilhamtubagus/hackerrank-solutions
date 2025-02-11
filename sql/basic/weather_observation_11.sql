// https://www.hackerrank.com/challenges/weather-observation-station-11

SELECT DISTINCT city
FROM station
WHERE city NOT REGEXP '[aeiouAEIOU]$' or city NOT REGEXP '^[aeiouAEIOU]'
