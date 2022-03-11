package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 指针思路
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}

	return pre
}

// 递归思路

func reverseList2(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}

	next := head.Next
	head.Next = pre
	return help(head, next)
}
