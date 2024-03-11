package alg

import "testing"

func TestCanJump(t *testing.T) {
	// Test case 1: Test with an empty nums slice
	nums1 := []int{}
	result1 := canJump(nums1)
	expected1 := false
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %t, got %t", expected1, result1)
	}

	// Test case 2: Test with a single element
	nums2 := []int{0}
	result2 := canJump(nums2)
	expected2 := true
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %t, got %t", expected2, result2)
	}

	// Test case 3: Test with non-jumpable numbers
	nums3 := []int{0, 2, 3, 0, 0}
	result3 := canJump(nums3)
	expected3 := false
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %t, got %t", expected3, result3)
	}

	// Test case 4: Test with jumpable numbers
	nums4 := []int{2, 3, 1, 1, 4}
	result4 := canJump(nums4)
	expected4 := true
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %t, got %t", expected4, result4)
	}

	// Test case 5: Test with non-jumpable numbers
	nums5 := []int{3, 2, 1, 0, 4}
	result5 := canJump(nums5)
	expected5 := false
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %t, got %t", expected5, result5)
	}

	// Add more test cases here...
}
