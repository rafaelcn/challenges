type NumArray struct {
    numbers []int
    prefix []int
}


func Constructor(nums []int) NumArray {
    a := NumArray{numbers: nums}
    a.prefix = make([]int, len(nums))
    
    copy(a.prefix, nums)

    for i := 1; i <= len(nums)-1; i++ {
        a.prefix[i] += a.prefix[i-1]
    }

    return a
}


func (n *NumArray) SumRange(left int, right int) int {
    if left > 0 {
        return n.prefix[right]-n.prefix[left-1]
    }

    return n.prefix[right]
}
