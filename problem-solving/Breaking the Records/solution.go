package solution

func BreakingRecords(scores []int32) []int32 {
	//lowest and highest value
	lowest := scores[0]
	highest := scores[0]
	//counter for lowest and highest value
	cL := 0
	cH := 0
	for _, score := range scores {
		if score > highest {
			highest = score
			cH++
		} else if score < lowest {
			lowest = score
			cL++
		}
	}
	return []int32{int32(cH), int32(cL)}

}
