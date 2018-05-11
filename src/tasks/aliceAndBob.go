package tasks

func AliceAndBob(a0 int32, a1 int32, a2 int32, b0 int32, b1 int32, b2 int32) []int32 {
	aliceIndex, bobIndex := 0, 1
	alice := [3]int32{a0, a1, a2}
	bob := [3]int32{b0, b1, b2}
	result := make([]int32, 2)
	for i := 0; i < len(alice); i++ {
		if alice[i] > bob[i] {
			result[aliceIndex]++
		} else if alice[i] < bob[i] {
			result[bobIndex]++
		}
	}
	return result
}
