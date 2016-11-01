package _02_add_two_numbers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T)  {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	answer := &ListNode{Val: 7, Next:&ListNode{Val:0, Next: &ListNode{Val:0, Next: &ListNode{Val: 1, Next:nil}}}}
	li := addTwoNumbers(l1, l2)
	assert.Equal(t, li, answer)
}

func BenchmarkAddTwoNumbers(t *testing.B)  {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	for i := 0; i < t.N; i++ {
		addTwoNumbers(l1, l2)
	}
}

