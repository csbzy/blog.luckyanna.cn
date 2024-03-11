package alg

// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  这个问题可以使用递归来解决。基本思路是这样的：

// 首先，我们需要检查链表中是否有足够的节点可以进行翻转。我们可以通过遍历链表来实现这一点。

// 如果有足够的节点，我们就进行翻转。我们可以通过递归调用来实现这一点。在每一次递归调用中，我们都会翻转一组节点，并将这一组的最后一个节点（也就是翻转后的第一个节点）连接到下一组翻转后的链表上。

// 如果没有足够的节点，我们就保持原有的顺序。

// 以下是解决这个问题的Go代码：

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}

	newHead, newTail := reverse(head, node)
	newTail.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) (*ListNode, *ListNode) {
	var (
		prev = last
		cur  = first
	)

	for cur != last {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev, first
}

// 	if head == nil {
// 		return nil
// 	}

// 	node := head
// 	for i := 0; i < k; i++ {
// 		if node == nil {
// 			return head
// 		}
// 		node = head.Next

// 	}
// 	fmt.Println(head, node)
// 	newHead, newTail := reverse(head, node)
// 	newTail.Next = reverseKGroup(node, k)
// 	return newHead
// }

// func reverse(first *ListNode, end *ListNode) (*ListNode, *ListNode) {
// 	prev := end
// 	cur := first
// 	fmt.Println("begin:", cur, end, first.Next)
// 	for cur != end {
// 		next := cur.Next
// 		cur.Next = prev
// 		prev = cur
// 		cur = next
// 	}
// 	fmt.Println(prev, first)
// 	return prev, first
// }

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	if head == nil {
// 		return head
// 	}

// 	node := head
// 	for i := 0; i < k; i++ {
// 		if node == nil {
// 			return head
// 		}
// 		node = node.Next
// 	}

// 	newHead, newTail := reverse(head, node)
// 	newTail.Next = reverseKGroup(node, k)
// 	return newHead
// }

// func reverse(first *ListNode, last *ListNode) (*ListNode, *ListNode) {
// 	var (
// 		prev = last
// 		cur  = first
// 	)

// 	for cur != last {
// 		next := cur.Next
// 		cur.Next = prev
// 		prev = cur
// 		cur = next
// 	}
// 	return prev, first
// }

func reverseKGroup1(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}

	newHead, newTail := reverse1(head, node)
	newTail.Next = reverseKGroup1(node, k)
	return newHead
}

func reverse1(first *ListNode, last *ListNode) (*ListNode, *ListNode) {
	prev := last
	cur := first
	for cur != last {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev, first
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseGroup(head *ListNode, k int) *ListNode {
// 	node := head
// 	for i := 0; i < k; i++ {
// 		if node == nil {
// 			return head
// 		}
// 		node = node.Next
// 	}

// 	newHead, newTail := reverse1(head, node)
// 	newTail.Next = reverseGroup(node, k)
// 	return newHead
// }

// func reverse1(first, last *ListNode) (*ListNode, *ListNode) {
// 	cur := first
// 	prev := last
// 	for cur != last {
// 		next := cur.Next
// 		cur.Next = prev
// 		prev = cur
// 		cur = next
// 	}

// 	return prev, first
// }
