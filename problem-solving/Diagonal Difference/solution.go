package solution

import (
	"math"
)

func DiagonalDifference(arr [][]int32) int32 {
	// Write your code here
	n := len(arr)
	diag := int32(0)
	antidiag := int32(0)
	for i := 0; i < n; i++ {
		diag += arr[i][i]
		antidiag += arr[(n-1)-i][i]
	}
	absVal := math.Abs(float64(diag) - float64(antidiag))
	return int32(absVal)
}
