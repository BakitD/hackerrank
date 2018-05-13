package implementation

func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var result int32
	for i := int32(0); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				result++
			}
		}
	}
	return result
}
