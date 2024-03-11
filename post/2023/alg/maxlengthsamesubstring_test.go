package alg

import "testing"

func TestMaxSameSubString(t *testing.T) {
	// Test case 1: Test with two strings having a common substring
	result1 := MaxSameSubString("abcdef", "defghi")
	expected1 := "def"
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %s, got %s", expected1, result1)
	}

	// Test case 2: Test with two strings having no common substring
	result2 := MaxSameSubString("abc", "def")
	expected2 := ""
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %s, got %s", expected2, result2)
	}

	// Test case 3: Test with two strings having multiple common substrings
	result3 := MaxSameSubString("abcde", "cdefg")
	expected3 := "cde"
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %s, got %s", expected3, result3)
	}

	// Add more test cases here...
}