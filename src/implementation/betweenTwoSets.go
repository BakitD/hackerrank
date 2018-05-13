package implementation

func BetweenTwoSets(a []int32, b []int32) int32 {
	factors := int32(0)
	for value := b[0]; value > 0; value-- {
		flag := true
		for _, aValue := range a {
			if value%aValue != 0 {
				flag = false
				break
			}
		}
		if flag {
			for _, bValue := range b {
				if bValue%value != 0 {
					flag = false
					break
				}
			}
		}
		if flag {
			factors++
		}
	}
	return factors
}
