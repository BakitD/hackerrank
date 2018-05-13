package implementation

func MigratoryBirds(ar []int32) int32 {
	var birds map[int32]int32 = map[int32]int32{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	for _, a := range ar {
		birds[a]++
	}
	var maximal, index int32 = 0, 0

	for i := int32(1); i <= 5; i++ {
		if birds[i] > maximal {
			maximal = birds[i]
			index = i
		}
	}
	return index
}
