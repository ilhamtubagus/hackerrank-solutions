package main

import solution "github.com/ilhamtubagus/hackerrank-solutions/problem-solving/go/solutions"

func main() {
	// // birtday cake
	// candles := []int32{1, 3, 1, 21, 30, 11, 30, 30, 21}
	// fmt.Println(solution.BirthdayCakeCandles(candles))

	// // diagonal difference
	// val := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	// fmt.Println(solution.DiagonalDifference(val))

	// // grading student
	// fmt.Println(solution.GradingStudents([]int32{38, 84, 73}))

	// // get movies title and sort hackerrank
	// c := solution.GetMovieTitles("spiderman")
	// for _, v := range c {
	// 	fmt.Println(v)
	// }
	// super stack hackerrank
	operation := []string{"push 1", "pop", "inc 5 2", "push -123", "inc 4 12"}
	solution.SuperStack(operation)
}
