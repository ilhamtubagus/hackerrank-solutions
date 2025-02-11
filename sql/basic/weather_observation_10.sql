// https://www.hackerrank.com/challenges/weather-observation-station-10

SELECT DISTINCT city
FROM station
WHERE city NOT REGEXP '[aeiouAEIOU]$'
