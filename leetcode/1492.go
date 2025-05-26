func kthFactor(n int, k int) int {
    l := []int{}
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            l = append(l, i)
        }
    }

    if len(l) < k {
        return -1
    }

    return l[k-1]
}
