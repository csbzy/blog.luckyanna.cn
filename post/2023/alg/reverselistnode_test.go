package alg

import (
	"reflect"
	"testing"
)

func TestReverseNode(t *testing.T) {
	// Test case 1: Reverse a list with multiple nodes
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	expected1 := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}
	result1 := reverseNode(head1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	}

	// Test case 2: Reverse a list with a single node
	head2 := &ListNode{Val: 1, Next: nil}
	expected2 := &ListNode{Val: 1, Next: nil}
	result2 := reverseNode(head2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	}

	// Test case 3: Reverse an empty list
	var head3 *ListNode
	expected3 := (*ListNode)(nil)
	result3 := reverseNode(head3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected3, result3)
	}

	// Add more test cases here...
}