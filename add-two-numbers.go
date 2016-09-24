package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	//l1 := &ListNode{Val: 1, Next: &ListNode{Val: 8, Next:nil}}
	//l2 := &ListNode{Val: 0, Next: nil}
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	li := addTwoNumbers(l1, l2)
	for l := li; l != nil ; l = l.Next {
		fmt.Println(l.Val)
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	carry := 0
	cl1, cl2, crl := l1, l2, result
	for {
		l1val := 0
		l2val := 0
		if cl1 != nil {
			l1val = cl1.Val
		}
		if cl2 != nil {
			l2val = cl2.Val
		}

		sum := l1val + l2val + carry
		crl.Val = sum % 10

		carry = sum / 10

		if cl1 != nil {
			cl1 = cl1.Next
		} else {
			cl1 = nil
		}
		if cl2 != nil {
			cl2 = cl2.Next
		} else {
			cl2 = nil
		}

		if cl1 == nil && cl2 == nil {
			break
		}
		crl.Next = &ListNode{}
		crl = crl.Next
	}

	if carry > 0 {
		crl.Next = &ListNode{Val: carry, Next: nil}
	}
	return result
}
