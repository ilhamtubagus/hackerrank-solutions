package main

import (
	"fmt"
	"math"
)

func lonelyinteger(a []int32) int32 {
	// Write your code here
	var mapValue map[int32]int32 = make(map[int32]int32, len(a))
	var lonelyint32 int32
	for _, val := range a {
		mapValue[val] += 1

		if mapValue[val] == 1 {
			lonelyint32 = val
		}
	}

	return lonelyint32
}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	pIndex := int32(0)
	sIndex := len(arr[0]) - 1
	pSum := int32(0)
	sSum := int32(0)
	for i := 0; i < len(arr); i++ {
		pSum += arr[i][pIndex]
		sSum += arr[i][sIndex]

		pIndex += 1
		sIndex -= 1
	}

	return int32(math.Abs(float64(pSum) - float64(sSum)))
}

func flippingMatrix(matrix [][]int) int {
	n := len(matrix) / 2
	sum := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			maxValue := max(
				matrix[i][j],
				matrix[i][2*n-j-1],
				matrix[2*n-i-1][j],
				matrix[2*n-i-1][2*n-j-1],
			)
			formatter := "Matrix [%v][%v] = %v\n"
			fmt.Printf(formatter, i, j, matrix[i][j])
			fmt.Printf(formatter, i, 2*n-j-1, matrix[i][2*n-j-1])
			fmt.Printf(formatter, 2*n-i-1, j, matrix[2*n-i-1][j])
			fmt.Printf(formatter, 2*n-i-1, 2*n-j-1, matrix[2*n-i-1][2*n-j-1])
			fmt.Println(maxValue)
			sum += maxValue
		}
	}
	return sum
}

func max(a, b, c, d int) int {
	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	if d > m {
		m = d
	}
	return m
}

func main() {
	matrix := [][]int{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108},
	}

	result := flippingMatrix(matrix)
	fmt.Println("Maximum Sum:", result)
}
