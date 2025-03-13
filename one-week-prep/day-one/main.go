package main

import (
	"fmt"
	"sort"
	"strconv"
)

func plusMinus(arr []int32) {
	// Write your code here
	var positive, negative, zero float64
	for _, val := range arr {
		if val > 0 {
			positive += 1
		} else if val < 0 {
			negative += 1
		} else {
			zero += 1
		}
	}
	length := float64(len(arr))
	fmt.Printf("%.6f\n", positive/length)
	fmt.Printf("%.6f\n", negative/length)
	fmt.Printf("%.6f\n", zero/length)
}

func timeConversion(s string) string {
	// Write your code here
	AMorPM := s[len(s)-2:]
	var result string
	if AMorPM == "AM" {
		hour := s[:2]
		rest := s[2:]
		if hour == "12" {
			hour = "00"
		}

		result = fmt.Sprintf("%v%v", hour, rest)
	} else if AMorPM == "PM" {
		hour := s[:2]
		rest := s[2:]
		if hour != "12" {
			hourInt, _ := strconv.Atoi(hour)
			hour24 := hourInt + 12
			hour = strconv.Itoa(hour24)
		}

		result = fmt.Sprintf("%v%v", hour, rest)
	}
	return result
}

func findMedian(arr []int32) int32 {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	length := len(arr)
	medianIndex := length / 2
	fmt.Println(length)
	fmt.Println(arr)

	return arr[medianIndex]
}

func main() {
	//plusMinus([]int32{0, 1, 0, -4, 6, 5, -3})
	//fmt.Println(timeConversion("07:05:45PM"))
	//fmt.Println(findMedian([]int32{0, 1, 2, 4, 6, 5, 3}))
	//fmt.Println(lonelyinteger([]int32{1, 1, 2, 2, 3, 2, 4, 3}))
	//fmt.Println(diagonalDifference([][]int32{
	//	{11, 2, 4},
	//	{4, 5, 6},
	//	{10, 8, -12},
	//}))
}
