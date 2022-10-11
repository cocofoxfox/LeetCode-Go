package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// define ListNode
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	carry, curNode := 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		d1, d2 := 0, 0
		if l1 != nil {
			d1, l1 = l1.Val, l1.Next
		}

		if l2 != nil {
			d2, l2 = l2.Val, l2.Next
		}

		curSum := d1 + d2 + carry
		curNode.Next = &ListNode{Val: curSum % 10}
		carry = curSum / 10
		curNode = curNode.Next
	}
	return head.Next
}
