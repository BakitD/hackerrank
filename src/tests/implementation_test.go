package tests

import (
	"implementation"
	"reflect"
	"testing"
)

func TestGradingStudents(t *testing.T) {
	testArray := []int32{73, 67, 38, 33}
	expectedArray := []int32{75, 67, 40, 33}
	result := implementation.GradingStudents(testArray)
	if !reflect.DeepEqual(result, expectedArray) {
		t.Errorf("Expected %v got %v", expectedArray, result)
	}
}

func ExampleApplesAndOranges() {
	implementation.CountApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
	// Output:
	// 1
	// 1
}
