package alg

import "testing"

func TestMaxDiv(t *testing.T) {
	// Test case 1: Test with two prime numbers
	result1 := MaxDiv(7, 11)
	expected1 := 1
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Test with two composite numbers
	result2 := MaxDiv(12, 18)
	expected2 := 6
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with one prime and one composite number
	result3 := MaxDiv(5, 15)
	expected3 := 5
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Add more test cases here...
}