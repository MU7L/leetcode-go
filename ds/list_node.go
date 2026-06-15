package ds

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(value []int) *ListNode {
	head := &ListNode{}
	cursor := head
	for _, val := range value {
		cursor.Next = &ListNode{
			Val: val,
		}
		cursor = cursor.Next
	}
	return head.Next
}
