package alg

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	// Test case 1: Test ongyulwith two empty lists
	l1 := createList([]int{})
	l2 := createList([]int{})
	result1 := addTwoNumbers(l1, l2)
	expected1 := createList([]int{})
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	}

	// Test case 2: Test with one empty list and one non-empty list
	l3 := createList([]int{})
	l4 := createList([]int{1, 2, 3})
	result2 := addTwoNumbers(l3, l4)
	expected2 := createList([]int{1, 2, 3})
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	}

	// Test case 3: Test with two non-empty lists of different lengths
	l5 := createList([]int{1, 2, 3})
	l6 := createList([]int{4, 5})
	result3 := addTwoNumbers(l5, l6)
	expected3 := createList([]int{5, 7, 3})
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected3, result3)
	}

	// Test case 4: Test with two non-empty lists of the same length
	l7 := createList([]int{1, 2, 3})
	l8 := createList([]int{4, 5, 6})
	result4 := addTwoNumbers(l7, l8)
	expected4 := createList([]int{5, 7, 9})
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Test case 4 failed: expected %v, got %v", expected4, result4)
	}

	// Add more test cases here...
}

func createList(nums []int) *ListNode {
	var head *ListNode
	var cur *ListNode
	for _, num := range nums {
		if cur == nil {
			cur = &ListNode{
				Val: num,
			}
			head = cur
		} else {
			cur.Next = &ListNode{
				Val: num,
			}
			cur = cur.Next
		}
	}
	return head
}