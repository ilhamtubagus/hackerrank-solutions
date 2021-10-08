package solution

import "sort"

func BirthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	sort.Slice(candles, func(i, j int) bool {
		return candles[i] > candles[j]
	})
	total := int32(1)
	prev := candles[0]
	for i, val := range candles {
		if i != 0 {
			if prev != val {
				break
			} else {
				total += 1
				prev = val
			}
		}
	}
	return total
}
