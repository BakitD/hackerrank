package implementation

//{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}

func BirthdayChocolate(n, m int32, array []int32) int32 {
	right := m
	arrayLength := int32(len(array))
	counts := int32(0)
	startSum := int32(0)
	for i := int32(0); i < right; i++ {
		startSum += array[i]
	}
	if startSum == n {
		counts++
	}
	for i := int32(right); i < arrayLength; i++ {
		startSum = startSum - array[i-m] + array[i]
		if startSum == n {
			counts++
		}
	}
	return counts
}
