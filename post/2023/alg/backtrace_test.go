package alg

import (
	"reflect"
	"testing"
)

func TestBacktrack(t *testing.T) {
	// Test case 1: Test with an empty nums slice
	nums1 := []int{}
	var result1 [][]int
	expected1 := [][]int{{}}
	backtrack(nums1, []int{}, &result1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	}

	// Test case 2: Test with a single element
	nums2 := []int{1}
	var result2 [][]int
	expected2 := [][]int{{1}}
	backtrack(nums2, []int{}, &result2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	}

	// Test case 3: Test with multiple elements
	nums3 := []int{1, 2, 3}
	var result3 [][]int
	expected3 := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	backtrack(nums3, []int{}, &result3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected3, result3)
	}

	// Add more test cases here...
}
