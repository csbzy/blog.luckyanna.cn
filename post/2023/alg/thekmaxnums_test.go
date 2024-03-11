package alg

import "testing"

func TestFindKthLargest(t *testing.T) {
	// Test case 1: Test with a sorted array in ascending order
	nums1 := []int{1, 2, 3, 4, 5}
	k1 := 3
	expected1 := 3
	result1 := findKthLargest(nums1, k1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Test with a sorted array in descending order
	nums2 := []int{5, 4, 3, 2, 1}
	k2 := 2
	expected2 := 4
	result2 := findKthLargest(nums2, k2)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with an unsorted array
	nums3 := []int{9, 7, 5, 3, 1}
	k3 := 4
	expected3 := 3
	result3 := findKthLargest(nums3, k3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Add more test cases here...

	// Test case 4: Test with an array containing duplicate elements
	nums4 := []int{2, 4, 6, 8, 8, 10}
	k4 := 1
	expected4 := 10
	result4 := findKthLargest(nums4, k4)
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Test case 5: Test with an empty array
	nums5 := []int{}
	k5 := 1
	expected5 := 0
	result5 := findKthLargest(nums5, k5)
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}

	// Test case 6: Test with an array containing negative numbers
	nums6 := []int{-5, -3, -1, -7, -9}
	k6 := 3
	expected6 := -5
	result6 := findKthLargest(nums6, k6)
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	}

	// Add more test cases here...
}
