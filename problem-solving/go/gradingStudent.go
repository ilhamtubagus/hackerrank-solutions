package main

func nextMltpl(n int32) int32 {
	return 5 * ((n / 5) + 1)
}

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	for i := 0; i < len(grades); i++ {
		if (nextMltpl(grades[i])-grades[i]) < 3 && grades[i] >= 38 {
			grades[i] = nextMltpl(grades[i])
		}
	}
	return grades
}
