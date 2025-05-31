func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }

    mid := len(lists)/2

    return  merge(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    l := head

    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            l.Next = l1
            l1 = l1.Next // walk over l1
        } else {
            l.Next = l2
            l2 = l2.Next // walk over l2
        }

        l = l.Next
    }

    for l1 != nil {
        l.Next = l1
        l = l.Next

        l1 = l1.Next
    }
    for l2 != nil {
        l.Next = l2
        l = l.Next

        l2 = l2.Next
    }

    return head.Next
}

func print(l *ListNode) string {
    s := "["
    for l != nil {
        s += fmt.Sprintf("%d,", l.Val)
        l = l.Next
    } 
    s += "]"

    return s
}
