package alg

import "testing"

func TestLongestConsecutive(t *testing.T) {
	// Test case 1: Test with an empty nums slice
	nums1 := []int{}
	result1 := longestConsecutive(nums1)
	expected1 := 0
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Test with consecutive numbers
	nums2 := []int{1, 2, 3, 4, 5}
	result2 := longestConsecutive(nums2)
	expected2 := 5
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with non-consecutive numbers
	nums3 := []int{1, 3, 5, 7, 9}
	result3 := longestConsecutive(nums3)
	expected3 := 1
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Test case 4: Test with duplicate numbers
	nums4 := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	result4 := longestConsecutive(nums4)
	expected4 := 4
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Add more test cases here...

	// Test case 5: Test with negative numbers
	nums5 := []int{-5, -4, -3, -2, -1}
	result5 := longestConsecutive(nums5)
	expected5 := 5
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}

	// Test case 6: Test with mixed positive and negative numbers
	nums6 := []int{-3, -2, -1, 0, 1, 2, 3}
	result6 := longestConsecutive(nums6)
	expected6 := 7
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	}

	// Test case 7: Test with a single element
	nums7 := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	result7 := longestConsecutive(nums7)
	expected7 := 7
	if result7 != expected7 {
		t.Errorf("Test case 7 failed: expected %d, got %d", expected7, result7)
	}

	// Add more test cases here...
}
