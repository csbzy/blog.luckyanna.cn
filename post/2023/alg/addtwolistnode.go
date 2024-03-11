package alg

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {


	var (
		head *ListNode
		cur  *ListNode
		in   = 0
	)
	for l1 != nil || l2 != nil {
		curSum := in
		if l2 != nil {
			curSum += l2.Val
			l2 = l2.Next
		}

		if l1 != nil {
			curSum += l1.Val
			l1 = l1.Next
		}

		curSum, in = check(curSum)
		if cur == nil {
			cur = &ListNode{
				Val: curSum,
			}
			head = cur
		} else {
			cur.Next = &ListNode{
				Val: curSum,
			}
			cur = cur.Next
		}

	}
	if in > 0 {
		cur.Next = &ListNode{
			Val: in,
		}
	}

	return head
}

func check(a int) (int, int) {
	if a >= 10 {
		return a - 10, 1
	}
	return a, 0
}

func reverseTwoNumber(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
