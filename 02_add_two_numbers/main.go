package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// var node = &ListNode{
	// 	Val: 0,
	// 	Next: &ListNode{
	// 		Val: 1,
	// 		Next: &ListNode{
	// 			Val: 2,
	// 			Next: &ListNode{
	// 				Val: 3,
	// 				Next: &ListNode{
	// 					Val: 4,
	// 					Next: &ListNode{
	// 						Val:  5,
	// 						Next: nil,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// fmt.Println(node)

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2))

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		retNode *ListNode
		carry   int
	)

	tmpL1, tmpL2 := l1, l2
	cursor := &retNode
	for tmpL1 != nil || tmpL2 != nil {

		val1 := getValueAndMoveNext(&tmpL1)
		val2 := getValueAndMoveNext(&tmpL2)

		*cursor = new(ListNode)
		(*cursor).Val = carry + val1 + val2
		carry = (*cursor).Val / 10
		(*cursor).Val %= 10

		cursor = &(*cursor).Next
	}

	if carry > 0 {
		*cursor = new(ListNode)
		(*cursor).Val = carry
	}

	return retNode
}

func getValueAndMoveNext(n **ListNode) int {
	if n == nil || *n == nil {
		return 0
	}
	ret := (*n).Val
	*n = (*n).Next
	return ret
}

func (node *ListNode) String() string {
	var s string
	for node != nil {
		s += string([]byte{byte(node.Val) + '0'})
		if node.Next != nil {
			s += " -> "
		}
		node = node.Next
	}
	return s
}
