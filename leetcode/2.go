/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    carry := 0

    head := &ListNode{Val:0,Next:nil}
    curr := head

    for l1 != nil || l2 != nil || carry != 0 {
        x := 0
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        }
        y := 0
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }

        s := carry + x + y
        carry = s/10

        curr.Next = &ListNode{Val: s%10}
        curr = curr.Next
    }

    return head.Next
}
