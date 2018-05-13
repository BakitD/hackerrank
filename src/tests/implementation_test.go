package tests

import (
	"fmt"
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

func ExampleWillKangarooMeetOne() {
	fmt.Print(implementation.WillKangarooMeet(0, 3, 4, 2))
	// Output:
	// YES

}

func ExampleWillKangarooMeetTwo() {

	fmt.Print(implementation.WillKangarooMeet(0, 2, 5, 3))
	// Output:
	// NO
}

func TestBetweenTwoSets(t *testing.T) {
	testA := []int32{2, 4}
	testB := []int32{16, 32, 96}
	expectedValue := int32(3)
	result := implementation.BetweenTwoSets(testA, testB)
	if result != expectedValue {
		t.Errorf("Expected %d got %d", expectedValue, result)
	}
}

func TestBreakingTheRecords(t *testing.T) {
	testArray := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	expectedArray := []int32{2, 4}
	result := implementation.BreakingRecords(testArray)
	if !reflect.DeepEqual(result, expectedArray) {
		t.Errorf("Expected %v got %v", expectedArray, result)
	}
}

func TestBirthdayChocolate(t *testing.T) {
	testArray := []int32{1, 2, 1, 3, 2}
	var n, m int32 = 3, 2
	result := implementation.BirthdayChocolate(n, m, testArray)
	expectedArray := int32(2)
	if result != expectedArray {
		t.Errorf("Expected %v got %v", expectedArray, result)
	}

	testArrayTwo := []int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}
	n, m = 18, 7
	result = implementation.BirthdayChocolate(n, m, testArrayTwo)
	expectedArray = int32(3)
	if result != expectedArray {
		t.Errorf("Expected %v got %v", expectedArray, result)
	}

}

func TestDivisibleSumPairs(t *testing.T) {
	testArray := []int32{1, 3, 2, 6, 1, 2}
	var n, k int32 = 6, 3
	var expected int32 = 5
	result := implementation.DivisibleSumPairs(n, k, testArray)
	if result != expected {
		t.Errorf("Expected %d got %d", expected, result)
	}
}

func TestMigratoryBirds(t *testing.T) {
	testArray := []int32{1, 4, 4, 4, 5, 3}
	expected := int32(4)
	result := implementation.MigratoryBirds(testArray)
	if result != expected {
		t.Errorf("Expected %d got %d", expected, result)
	}
}
