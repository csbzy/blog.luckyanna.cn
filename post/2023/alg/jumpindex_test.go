package alg

import "testing"

func TestJump(t *testing.T) {
	// // Test case 1: Test with an empty nums slice
	// nums1 := []int{}
	// result1 := jump(nums1)
	// expected1 := 0
	// if result1 != expected1 {
	// 	t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	// }

	// // Test case 2: Test with a single element
	// nums2 := []int{5}
	// result2 := jump(nums2)
	// expected2 := 0
	// if result2 != expected2 {
	// 	t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	// }

	// // Test case 3: Test with consecutive numbers
	// nums3 := []int{1, 2, 3, 4, 5}
	// result3 := jump(nums3)
	// expected3 := 4
	// if result3 != expected3 {
	// 	t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	// }

	// // Test case 4: Test with non-consecutive numbers
	// nums4 := []int{1, 3, 5, 7, 9}
	// result4 := jump(nums4)
	// expected4 := 4
	// if result4 != expected4 {
	// 	t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	// }

	// // Test case 5: Test with duplicate numbers
	// nums5 := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	// result5 := jump(nums5)
	// expected5 := 4
	// if result5 != expected5 {
	// 	t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	// }

	// // Add more test cases here...

	// // Test case 6: Test with negative numbers
	// nums6 := []int{-5, -4, -3, -2, -1}
	// result6 := jump(nums6)
	// expected6 := 4
	// if result6 != expected6 {
	// 	t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	// }

	// Test case 7: Test with mixed positive and negative numbers
	nums7 := []int{2, 3, 1}
	result7 := jump(nums7)
	expected7 := 1
	if result7 != expected7 {
		t.Errorf("Test case 7 %v failed: expected %d, got %d", nums7, expected7, result7)
	}

	// Test case 8: Test with large numbers
	nums8 := []int{2, 3, 1, 1, 4}
	result8 := jump(nums8)
	expected8 := 2
	if result8 != expected8 {
		t.Errorf("Test case 8 failed: expected %d, got %d", expected8, result8)
	}

	// Add more test cases here...
	// Test case 8: Test with large numbers
	nums9 := []int{2, 3, 3, 1, 1, 1, 1, 4}
	result9 := jump(nums9)
	expected9 := 4
	if result9 != expected9 {
		t.Errorf("Test case 9 failed: expected %d, got %d", expected9, result9)
	}
}
