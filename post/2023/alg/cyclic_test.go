package alg

import "testing"

func TestIsCyclic(t *testing.T) {
	// Test case 1: Test with an empty nodes slice
	nodes1 := []*Node{}
	result1 := isCyclic(nodes1)
	expected1 := false
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %t, got %t", expected1, result1)
	}

	// Test case 2: Test with a single node
	node2 := &Node{}
	nodes2 := []*Node{node2}
	result2 := isCyclic(nodes2)
	expected2 := false
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %t, got %t", expected2, result2)
	}

	// Test case 3: Test with acyclic nodes
	node31 := &Node{}
	node32 := &Node{}
	node33 := &Node{}
	nodes3 := []*Node{node31, node32, node33}
	result3 := isCyclic(nodes3)
	expected3 := false
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %t, got %t", expected3, result3)
	}

	// Test case 4: Test with cyclic nodes
	node41 := &Node{}
	node42 := &Node{}
	node43 := &Node{}
	node41.Children = append(node41.Children, node42)
	node42.Children = append(node42.Children, node43)
	node43.Children = append(node43.Children, node41)
	nodes4 := []*Node{node41, node42, node43}
	result4 := isCyclic(nodes4)
	expected4 := true
	if result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %t, got %t", expected4, result4)
	}

	// Add more test cases here...
}
