package alg

import "testing"

func TestMaxProfit(t *testing.T) {
	// Test case 1: Test with an empty prices slice
	prices1 := []int{}
	result1 := maxProfit(prices1)
	expected1 := 0
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Test with prices in ascending order
	prices2 := []int{1, 2, 3, 4, 5}
	result2 := maxProfit(prices2)
	expected2 := 4
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with prices in descending order
	prices3 := []int{5, 4, 3, 2, 1}
	result3 := maxProfit(prices3)
	expected3 := 0
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Test case 4: Test with prices in random order
	prices4 := []int{7, 1, 5, 3, 6, 4}
	result4 := maxProfit(prices4)
	expected4 := 5
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Add more test cases here...

	// Test case 5: Test with prices containing negative values
	prices5 := []int{2, -1, 4, -3, 5}
	result5 := maxProfit(prices5)
	expected5 := 8
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}

	// Test case 6: Test with prices containing duplicate values
	prices6 := []int{2, 2, 2, 2, 2}
	result6 := maxProfit(prices6)
	expected6 := 0
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	}

	// Test case 7: Test with prices containing only one element
	prices7 := []int{10}
	result7 := maxProfit(prices7)
	expected7 := 0
	if result7 != expected7 {
		t.Errorf("Test case 7 failed: expected %d, got %d", expected7, result7)
	}

	// Add more test cases here...
}
