package warmup

func BirthdayCakeCandles(n int32, ar []int32) (count int32) {
	maximal := int32(0)
	for _, value := range ar {
		if value > maximal {
			maximal = value
			count = 1
		} else if value == maximal {
			count++
		}
	}
	return count
}
