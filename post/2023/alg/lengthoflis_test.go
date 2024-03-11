package alg

import "testing"

func TestLengthOfLIS(t *testing.T) {
	// Test case 1: Test with an empty slice
	result1 := lengthOfLIS([]int{})
	expected1 := 0
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2:  with a slice containing one element
	result2 := lengthOfLIS([]int{5})
	expected2 := 1
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with a slice containing multiple elements
	result3 := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	expected3 := 4
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Add more test cases here...

	// Test case 4: Test with a slice containing all increasing elements
	result4 := lengthOfLIS([]int{1, 2, 3, 4, 5})
	expected4 := 5
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Test case 5: Test with a slice containing all decreasing elements
	result5 := lengthOfLIS([]int{5, 4, 3, 2, 1})
	expected5 := 1
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}

	// Test case 6: Test with a slice containing repeated elements
	result6 := lengthOfLIS([]int{1, 2, 2, 3, 4, 4, 5})
	expected6 := 5
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	}

	// Add more test cases here...
}
