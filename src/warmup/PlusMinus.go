package warmup

import "fmt"

func PlusMinus(arr []int32) {
	n := len(arr)
	pos, neg, neut := 0.0, 0.0, 0.0
	for i := 0; i < n; i++ {
		if arr[i] < 0 {
			neg++
		} else if arr[i] == 0 {
			neut++
		} else {
			pos++
		}
	}
	floatN := float64(n)
	fmt.Printf("%.6f\n%.6f\n%.6f", pos/floatN, neg/floatN, neut/floatN)
}
