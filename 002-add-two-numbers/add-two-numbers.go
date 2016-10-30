package _02_add_two_numbers

type ListNode struct {
	Val int
	Next *ListNode
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
