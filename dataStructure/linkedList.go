package dataStructure

import (
	"fmt"
	"math"
)

type ListNode struct {
	Data int
	Next *ListNode
}

type BiLinkedListNode struct {
	Data   int
	Next   *BiLinkedListNode
	Before *BiLinkedListNode
}

func PrintLinkedList(head *ListNode) {
	// Empty list check
	if head == nil {
		fmt.Println("Empty List")
		return
	}
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func ConvertSlice2LinkedList(slc []int) *ListNode {
	if len(slc) == 0 {
		return nil
	}
	head := &ListNode{Data: slc[0], Next: nil}
	current := head
	for i := 1; i < len(slc); i++ {
		current.Next = &ListNode{Data: slc[i], Next: nil}
		current = current.Next
	}
	return head
}

func CountLinkedList(head *ListNode) int {
	if head == nil {
		return 0
	}
	numOfNode := 0
	current := head
	for current != nil {
		numOfNode++
		current = current.Next
	}
	return numOfNode
}

func AppendNewNode(head *ListNode, data int) *ListNode {
	if head == nil {
		head = &ListNode{Data: data, Next: nil}
		return head
	}

	// Find the rear node
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Data: data, Next: nil}
	return head
}

func InsertNewNodeHead(head *ListNode, data int) *ListNode {
	if head == nil {
		head = &ListNode{Data: data, Next: nil}
		return head
	}
	newNode := &ListNode{Data: data, Next: nil}
	newNode.Next = head
	head = newNode
	return head
}

func InsertNewNode(head *ListNode, data int, position int) *ListNode {
	newNode := &ListNode{Data: data, Next: nil}
	if head == nil {
		return newNode
	}
	if position == 0 {
		return InsertNewNodeHead(head, data)
	}
	if position > CountLinkedList(head) {
		return AppendNewNode(head, data)
	}
	before := head
	traverseTimes := 0
	for traverseTimes < position-1 {
		before = before.Next
		traverseTimes++
	}
	after := before.Next
	before.Next = newNode
	newNode.Next = after
	return head
}

func DeleteNodeAtPosition(head *ListNode, position int) *ListNode {
	if position == 0 {
		return head.Next
	}

	current := head
	for i := 0; i < position-1; i++ {
		current = current.Next
	} // current is the previous node of deleted node
	succ := current.Next.Next
	current.Next = succ
	return head
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	arr := make([]int, 0)
	for current := head; current != nil; current = current.Next {
		arr = append(arr, current.Data)
	}
	reverseArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		reverseArr[len(arr)-i-1] = arr[i]
	}
	return ConvertSlice2LinkedList(reverseArr)
}

func traverseLinkedList(head *ListNode) []*ListNode {
	res := make([]*ListNode, 0)
	if head == nil {
		return res
	}
	current := head
	for current != nil {
		res = append(res, current)
		current = current.Next
	}
	return res
}

// GetIntersectionNode Leetcode 160
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	valA := traverseLinkedList(headA)
	valB := traverseLinkedList(headB)
	if valA[len(valA)-1] != valB[len(valB)-1] {
		return nil
	}
	var sameNode *ListNode
	for i := 0; i < int(math.Min(float64(len(valA)), float64(len(valB)))); i++ {
		idxA := len(valA) - i - 1
		idxB := len(valB) - i - 1
		if valA[idxA] == valB[idxB] {
			sameNode = valA[idxA]
		}
	}
	return sameNode
}
