package implementation

func WillKangarooMeet(x1, v1, x2, v2 int32) string {
	yes, no := "YES", "NO"
	if x1 < x2 {
		x1, x2 = x2, x1
		v1, v2 = v2, v1
	}
	flag := true
	diff := x1 - x2
	for step := int32(1); flag; step++ {
		x1 += v1
		x2 += v2
		newDiff := x1 - x2
		if x1 == x2 {
			break
		}
		if newDiff >= diff || x2 > x1 {
			return no
		}
		diff = newDiff
	}
	return yes
}
