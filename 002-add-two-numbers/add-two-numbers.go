package _02_add_two_numbers


type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tracer := l1
	carry, sum := 0, 0
	for  {
		sum = carry + tracer.Val + l2.Val
		carry = sum / 10
		tracer.Val = sum - carry * 10
		if tracer.Next == nil || l2.Next == nil{
			if tracer.Next == nil {
				tracer.Next = l2.Next
			}
			for carry > 0 {
				if tracer.Next == nil {
					tracer.Next = &ListNode{Val: carry, Next: nil}
					break
				}
				tracer = tracer.Next
				sum = carry + tracer.Val
				carry = sum / 10
				tracer.Val = sum - carry * 10
			}
			break
		}
		tracer, l2 = tracer.Next, l2.Next
	}

	return l1
}
