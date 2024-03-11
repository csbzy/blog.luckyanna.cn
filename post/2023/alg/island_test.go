package alg

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	// Test case 1: Test with an empty grid
	grid1 := [][]byte{}
	result1 := numIslands(grid1)
	expected1 := 0
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Test with a grid containing no islands
	grid2 := [][]byte{
		{'0', '0', '0'},
		{'0', '0', '0'},
		{'0', '0', '0'},
	}
	result2 := numIslands(grid2)
	expected2 := 0
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: Test with a grid containing one island
	grid3 := [][]byte{
		{'1', '1', '1'},
		{'1', '1', '1'},
		{'1', '1', '1'},
	}
	result3 := numIslands(grid3)
	expected3 := 1
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Add more test cases here...

	// Test case 4: Test with a grid containing multiple islands
	grid4 := [][]byte{
		{'1', '0', '1'},
		{'0', '1', '0'},
		{'1', '0', '1'},
	}
	result4 := numIslands(grid4)
	expected4 := 5
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}

	// Test case 5: Test with a grid containing a single row
	grid5 := [][]byte{
		{'1', '1', '1', '1'},
	}
	result5 := numIslands(grid5)
	expected5 := 1
	if result5 != expected5 {
		t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
	}

	// Test case 6: Test with a grid containing a single column
	grid6 := [][]byte{
		{'1'},
		{'1'},
		{'1'},
		{'1'},
	}
	result6 := numIslands(grid6)
	expected6 := 1
	if result6 != expected6 {
		t.Errorf("Test case 6 failed: expected %d, got %d", expected6, result6)
	}

	// Add more test cases here...
}
