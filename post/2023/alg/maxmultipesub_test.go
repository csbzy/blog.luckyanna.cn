package alg

import "testing"

func TestMaxProduct(t *testing.T) {
	// Test case 1: Test with an empty nums slice
	nums1 := []int{-2, 0, -1}
	result1 := maxProduct(nums1)
	expected1 := 0
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Test with a single element
	nums2 := []int{-1, -2, -3, 0}
	result2 := maxProduct(nums2)
	expected2 := 6
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with two elements
	nums3 := []int{2, 3}
	result3 := maxProduct(nums3)
	expected3 := 6
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Test case 4: Test with positive numbers
	nums4 := []int{2, -5, -2, -4, 3}
	result4 := maxProduct(nums4)
	expected4 := 24
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Test case 5: Test with negative numbers
	nums5 := []int{-2, -3, -4, -5}
	result5 := maxProduct(nums5)
	expected5 := 120
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}

	// Test case 6: Test with a mix of positive and negative numbers
	nums6 := []int{-2, 3, -4, 5}
	result6 := maxProduct(nums6)
	expected6 := 120
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	}

	// Add more test cases here...
}
