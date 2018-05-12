package tests

import (
	"reflect"
	"testing"
	"warmup"
)

func TestWarmUp(t *testing.T) {
	const ArraySize = 6
	testArray := [ArraySize]int32{1, 2, 3, 4, 10, 11}
	if warmup.SimpleArraySum(testArray[0:]) != 31 {
		t.Error("Expected 31")
	}
}

func TestDiagonalDifference(t *testing.T) {
	testArray := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	if warmup.DiagonalDifference(testArray) != 15 {
		t.Error("Expected 15")
	}
}

func TestAliceAndBob(t *testing.T) {
	expected := [2]int32{1, 1}
	if reflect.DeepEqual(warmup.AliceAndBob(5, 6, 7, 3, 6, 9), expected) {
		t.Error("Expected [1, 1]")
	}
}

func TestAVeryBigSum(t *testing.T) {
	testArray := [5]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	if warmup.AVeryBigSum(5, testArray[0:]) != 5000000015 {
		t.Error("Expected 5000000015")
	}
}

func ExamplePlusMinus() {
	warmup.PlusMinus([]int32{-4, 3, -9, 0, 4, 1})
	// Output:
	// 0.500000
	// 0.333333
	// 0.166667
}

func ExampleStaircase() {
	warmup.Staircase(4)
	// Output:
	//    #
	//   ##
	//  ###
	// ####
}

func ExampleMiniMaxSun() {
	warmup.MiniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
	// Output:
	// 2063136757 2744467344
}

func TestBirthdayCakeCandles(t *testing.T) {
	testArrayOne := []int32{3, 2, 1, 3}
	if warmup.BirthdayCakeCandles(4, testArrayOne[0:]) != 2 {
		t.Error("Expected 2")
	}
	testArrayTwo := []int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
	result := warmup.BirthdayCakeCandles(10, testArrayTwo[0:])
	if result != 5 {
		t.Error("Expected 5 got ", result)
	}
}

func TestTimeConversion(t *testing.T) {
	var testCases map[string]string = map[string]string{
		"01:00:00AM": "01:00:00",
		"07:05:05AM": "07:05:05",
		"07:05:05PM": "19:05:05",
		"11:00:00PM": "23:00:00",
		"12:00:05AM": "00:00:05",
		"12:00:05PM": "12:00:05",
		"00:00:00AM": "00:00:00",
	}
	for key, value := range testCases {
		result := warmup.TimeConversion(key)
		if value != result {
			t.Errorf("Expected %s got %s", value, result)
		}
	}
}
