package alg

import "testing"

func TestWordBreak(t *testing.T) {
	// Test case 1: Test with an empty string and empty wordDict
	s1 := ""
	wordDict1 := []string{}
	result1 := wordBreak(s1, wordDict1)
	expected1 := true
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	}

	// Test case 2: Test with a string that can be broken into words
	s2 := "aaaaaaa"
	wordDict2 := []string{"aaaa", "aaa"}
	result2 := wordBreak(s2, wordDict2)
	expected2 := true
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	}

	// Test case 3: Test with a string that cannot be broken into words
	s3 := "bb"
	wordDict3 := []string{"a", "b", "bbb", "bbbb"}
	result3 := wordBreak(s3, wordDict3)
	expected3 := true
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected3, result3)
	}

	// Test case 4: Test with a string that can be broken into words with duplicates
	s4 := "catsandog"
	wordDict4 := []string{"cats", "dog", "sand", "and", "cat"}
	result4 := wordBreak(s4, wordDict4)
	expected4 := false
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %v, got %v", expected4, result4)
	}

	// Add more test cases here...

	// Test case 5: Test with a string that can be broken into words with multiple matches
	s5 := "applepenapplepen"
	wordDict5 := []string{"apple", "pen"}
	result5 := wordBreak(s5, wordDict5)
	expected5 := true
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %v, got %v", expected5, result5)
	}

	// Test case 6: Test with a string that cannot be broken into words with empty wordDict
	s6 := "leetcode"
	wordDict6 := []string{}
	result6 := wordBreak(s6, wordDict6)
	expected6 := false
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %v, got %v", expected6, result6)
	}

	// Test case 7: Test with a string that can be broken into words with single character words
	s7 := "abcde"
	wordDict7 := []string{"a", "b", "c", "d", "e"}
	result7 := wordBreak(s7, wordDict7)
	expected7 := true
	if result7 != expected7 {
		t.Errorf("Test case 7 failed: expected %v, got %v", expected7, result7)
	}

	// Add more test cases here...
}
