package main

import (
	"fmt"
	"sort"
)

func isOdd(v int) bool {
	return v%2 != 0
}

/*
 * Complete the 'customSorting' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.
 */

func customSorting(strArr []string) []string {
	sort.Slice(strArr, func(i, j int) bool {
		x := len(strArr[i])
		y := len(strArr[j])
		if (isOdd(x) && !isOdd(y)) || (!isOdd(x) && isOdd(y)) {
			return isOdd(x)
		}

		if isOdd(x) && isOdd(y) {
			return x < y
		}

		if !isOdd(x) && !isOdd(y) {
			return x > y
		}

		if x == y {
			return strArr[i] < strArr[j]
		} else {
			return false
		}
	})

	return strArr
}

func main() {
	result := customSorting([]string{
		"abc",
		"ab",
		"abcde",
		"a",
		"abcd",
	})

	fmt.Println(result)
}
