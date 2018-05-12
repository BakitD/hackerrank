package warmup

func AVeryBigSum(n int32, ar []int64) int64 {
	var result int64
	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}
	return result
}
