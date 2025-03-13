package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * Complete the 'gridChallenge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func gridChallenge(grid []string) string {
	// Write your code here
	var gridString [][]string = make([][]string, len(grid))
	for i, val := range grid {
		row := strings.Split(val, "")
		sort.Slice(row, func(i, j int) bool {
			return row[i] < row[j]
		})
		gridString[i] = row
	}

	fmt.Println(gridString)
	for i := 0; i < len(gridString[0]); i++ {
		for j := 1; j < len(gridString); j++ {
			fmt.Printf("%v >< %v\n", gridString[j-1][i], gridString[j][i])
			if gridString[j-1][i] > gridString[j][i] {
				return "NO"
			}
		}
	}

	return "YES"
}

func main() {
	fmt.Println(gridChallenge([]string{"abc", "hjk", "mpq", "rtv"}))
}
