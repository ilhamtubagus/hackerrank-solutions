package Array_Manipulation

// https://www.hackerrank.com/challenges/crush/problem

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func arrayManipulation(n int32, queries [][]int32) int64 {
	// Create an array to store the differences
	arr := make([]int64, n+1)

	// Mark the start and end of each range
	for _, query := range queries {
		a := query[0] - 1
		b := query[1]
		k := int64(query[2])

		arr[a] += k
		if b < n {
			arr[b] -= k
		}
	}

	// Calculate prefix sums and find the maximum value
	max := int64(0)
	current := int64(0)
	for i := int32(0); i < n; i++ {
		current += arr[i]
		if current > max {
			max = current
		}
	}

	return max
}
