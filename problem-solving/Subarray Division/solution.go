package solution

// https://www.hackerrank.com/challenges/the-birthday-bar/problem
func sum(v []int32) int32 {
	var result int32
	for _, val := range v {
		result += val
	}
	return result
}
func Birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	c := 0
	for i := range s {
    //make sure that we are not accessing slice out of its capacity
		if i+int(m) >= len(s) {
			break
		}
		if sum(s[i:i+int(m)]) == d {
			c++
		}
	}
	return int32(c)
}

