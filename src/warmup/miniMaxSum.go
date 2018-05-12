package warmup

import (
	"fmt"
	"sort"
)

type Integer []int32

func (array Integer) Len() int           { return len(array) }
func (array Integer) Less(i, j int) bool { return array[i] < array[j] }
func (array Integer) Swap(i, j int)      { array[i], array[j] = array[j], array[i] }

func MiniMaxSum(arr []int32) {
	n := len(arr)
	toSort := make(Integer, n)
	for index, value := range arr {
		toSort[index] = value
	}
	sort.Sort(toSort[0:])
	var max, min int64
	for i := 0; i < n-1; i++ {
		max += int64(toSort[n-i-1])
		min += int64(toSort[i])
	}
	fmt.Println(min, max)
}
