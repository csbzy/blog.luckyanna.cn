package alg

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	// // Test case 1: k is greater than the length of the list
	// head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	// k1 := 4
	// expected1 := CopyList(head1)                  // Create a new copy of the linked list
	// result1 := reverseKGroup(CopyList(head1), k1) // Use the copied list for reverseKGroup
	// if !isEqual(result1, expected1) {
	// 	t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	// }

	// // Test case 2: k is equal to the length of the list
	// head2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	// k2 := 2
	// expected2 := reverseList(CopyList(head2))     // Create a new copy of the linked list
	// result2 := reverseKGroup(CopyList(head2), k2) // Use the copied list for reverseKGroup
	// if !isEqual(result2, expected2) {
	// 	t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	// }
	// head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	// newHead := reverseKGroup(head, 2)
	// // Print the reversed list
	// for node := newHead; node != nil; node = node.Next {
	// 	fmt.Println(node.Val)
	// }
	// // Test case 3: k is less than the length of the list
	// head3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	// k3 := 2
	// expected3 := &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}}
	// result3 := reverseKGroup(CopyList(head3), k3) // Use the copied list for reverseKGroup
	// if !isEqual(result3, expected3) {
	// 	t.Errorf("Test case 3 failed: expected %v, got %v", expected3, result3)
	// }

	// Add more test cases here...
}

// Helper function to create a copy of a linked list
func CopyList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Val: head.Val}
	current := newHead
	for head.Next != nil {
		head = head.Next
		current.Next = &ListNode{Val: head.Val}
		current = current.Next
	}

	return newHead
}

// Helper function to check if two linked lists are equal
func isEqual(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// Helper function to reverse a linked list
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
