package implementation

func BreakingRecords(score []int32) []int32 {
	min, max := score[0], score[0]
	mins, maxs := int32(0), int32(0)
	for _, value := range score {
		if value < min {
			min = value
			mins++
		}
		if value > max {
			max = value
			maxs++
		}
	}
	return []int32{maxs, mins}
}
