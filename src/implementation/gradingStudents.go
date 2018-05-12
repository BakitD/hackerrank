package implementation

func GradingStudents(grades []int32) []int32 {
	result := make([]int32, len(grades))
	for index, grade := range grades {
		value := grade
		if grade < 38 {
			result[index] = grade
			continue
		}
		for value = grade; value%5 != 0; value++ {
		}
		if value-grade < 3 {
			result[index] = value
		} else {
			result[index] = grade
		}
	}
	return result
}
