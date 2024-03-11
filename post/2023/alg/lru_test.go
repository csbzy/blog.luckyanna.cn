package alg

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	// Test case 1: Test with capacity 0
	// cache1 := New(0)
	// cache1.Put(1, -1)
	// if cache1.Get(1) != nil {
	// 	t.Errorf("Test case 1 failed: expected -1, got %d", cache1.Get(1))
	// }

	// Test case 2: Test with capacity 1
	cache2 := Constructor(1)
	cache2.Put(1, 1)
	if cache2.Get(1) != 1 {
		t.Errorf("Test case 2 failed: expected 1, got %d", cache2.Get(1))
	}
	cache2.Put(2, 2)
	if cache2.Get(1) != -1 {
		t.Errorf("Test case 2 failed: expected nil, got %d", cache2.Get(1))
	}
	if cache2.Get(2) != 2 {
		t.Errorf("Test case 2 failed: expected 2, got %d", cache2.Get(2))
	}

	// Test case 3: Test with capacity 2
	cache3 := Constructor(2)
	cache3.Put(1, 1)
	cache3.Put(2, 2)
	if cache3.Get(1) != 1 {
		t.Errorf("Test case 3 failed: expected 1, got %d", cache3.Get(1))
	}
	cache3.Put(3, 3)
	if cache3.Get(2) != -1 {
		t.Errorf("Test case 3 failed: expected nil, got %v", cache3.Get(2))
	}
	if cache3.Get(3) != 3 {
		t.Errorf("Test case 3 failed: expected 3, got %d", cache3.Get(3))
	}

	// Add more test cases here...
}
