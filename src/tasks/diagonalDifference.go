package tasks

import (
	"math"
)

func DiagonalDifference(a [][]int32) int32 {
	var diag, invDiag int32 = 0, 0
	n := len(a)
	for i := 0; i < n; i++ {
		diag += a[i][i]
		invDiag += a[i][n-i-1]
	}
	return int32(math.Abs(float64(diag - invDiag)))
}
