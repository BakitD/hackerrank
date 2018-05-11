package tests

import (
	"reflect"
	"tasks"
	"testing"
)

func TestWarmUp(t *testing.T) {
	const ArraySize = 6
	testArray := [ArraySize]int32{1, 2, 3, 4, 10, 11}
	if tasks.SimpleArraySum(testArray[0:]) != 31 {
		t.Error("Expected 31")
	}
}

func TestDiagonalDifference(t *testing.T) {
	testArray := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	if tasks.DiagonalDifference(testArray) != 15 {
		t.Error("Expected 15")
	}
}

func TestAliceAndBob(t *testing.T) {
	expected := [2]int32{1, 1}
	if reflect.DeepEqual(tasks.AliceAndBob(5, 6, 7, 3, 6, 9), expected) {
		t.Error("Expected [1, 1]")
	}
}

func TestAVeryBigSum(t *testing.T) {
	testArray := [5]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	if tasks.AVeryBigSum(5, testArray[0:]) != 5000000015 {
		t.Error("Expected 5000000015")
	}
}