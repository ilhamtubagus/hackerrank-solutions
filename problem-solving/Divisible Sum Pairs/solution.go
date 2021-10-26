package solution

func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var count int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}
