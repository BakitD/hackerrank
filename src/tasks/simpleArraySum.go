package tasks

func SimpleArraySum(ar []int32) int32 {
	result := int32(0)
	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}
	return result
}
