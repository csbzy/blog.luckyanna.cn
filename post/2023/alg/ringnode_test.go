package alg

import "testing"

func TestBackToOrigin(t *testing.T) {
	// Test case 1: Test with n = 0
	result1 := backToOrigin(0)
	expected1 := 1
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Test with n = 1
	result2 := backToOrigin(1)
	expected2 := 0
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with n = 2
	result3 := backToOrigin(2)
	expected3 := 2
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Test case 4: Test with n = 5
	result4 := backToOrigin(5)
	expected4 := 100
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Add more test cases here...

	// Test case 5: Test with n = 10
	result5 := backToOrigin(10)
	expected5 := 486
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}

	// Test case 6: Test with n = 15
	result6 := backToOrigin(15)
	expected6 := 3003
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	}

	// Test case 7: Test with n = 20
	result7 := backToOrigin(20)
	expected7 := 18476
	if result7 != expected7 {
		t.Errorf("Test case 7 failed: expected %d, got %d", expected7, result7)
	}

	// Add more test cases here...
}
