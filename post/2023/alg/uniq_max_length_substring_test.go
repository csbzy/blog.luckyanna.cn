package alg

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	// Test case 1: Test with a string containing unique characters
	result1 := lengthOfLongestSubstring("abcde")
	expected1 := 5
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Test with a string containing repeated characters
	result2 := lengthOfLongestSubstring("aabbc")
	expected2 := 2
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with an empty string
	result3 := lengthOfLongestSubstring("")
	expected3 := 0
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Add more test cases here...

	// Test case 4: Test with a string containing all unique characters
	result4 := lengthOfLongestSubstring("xyz")
	expected4 := 3
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Test case 5: Test with a string containing repeated characters at the beginning
	result5 := lengthOfLongestSubstring("aabbcc")
	expected5 := 2
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}

	// Test case 6: Test with a string containing repeated characters at the end
	result6 := lengthOfLongestSubstring("defgg")
	expected6 := 4
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	}

	// Add more test cases here...
}
