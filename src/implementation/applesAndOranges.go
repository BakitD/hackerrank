package implementation

import "fmt"

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	applesIn, orangesIn := 0, 0
	for _, position := range apples {
		if a+position >= s && a+position <= t {
			applesIn++
		}
	}
	for _, position := range oranges {
		if b+position >= s && b+position <= t {
			orangesIn++
		}
	}
	fmt.Printf("%d\n%d", applesIn, orangesIn)
}
