package alg

import "testing"

func TestTrap(t *testing.T) {
	// Test case 1: Test with an empty height array
	height1 := []int{}
	result1 := trap(height1)
	expected1 := 0
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Test with a height array containing no trapped water
	height2 := []int{1, 2, 3, 4, 5}
	result2 := trap(height2)
	expected2 := 0
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with a height array containing trapped water
	height3 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result3 := trap(height3)
	expected3 := 6
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Add more test cases here...

	// Test case 4: Test with a height array containing no trapped water at the edges
	height4 := []int{1, 0, 0, 0, 1}
	result4 := trap(height4)
	expected4 := 3
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Test case 5: Test with a height array containing trapped water at the edges
	height5 := []int{1, 0, 2, 0, 1}
	result5 := trap(height5)
	expected5 := 2
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}

	// Test case 6: Test with a height array containing equal heights
	height6 := []int{2, 2, 2, 2, 2}
	result6 := trap(height6)
	expected6 := 0
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	}

	height7 := []int{4, 2, 0, 3, 2, 5}
	result7 := trap(height7)
	expected7 := 9
	if result7 != expected7 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected7, result7)
	}

	// Add more test cases here...
}
