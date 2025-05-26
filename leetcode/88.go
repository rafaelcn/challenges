func merge(a []int, m int, b []int, n int)  {
    r  := []int{}

    i := 0
    j := 0

    for ; i < m && j < n; {
        if a[i] < b[j] {
            r = append(r, a[i])
            i++
        } else {
            r = append(r, b[j])
            j++
        }
    }

    for i < m {
        r = append(r, a[i])
        i++
    }

    for j < n {
        r = append(r, b[j])
        j++
    }

    for i := 0; i < len(a); i++ {
        a[i] = r[i]
    }
}
