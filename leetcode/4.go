func median(a []int) float64 {
    s := len(a)
    h := int(math.Floor(float64(s)/2))

    if s%2 == 0 {
        return float64(a[h]+a[h-1])/2
    }

    return float64(a[h])
}

func merge(a, b []int) []int {

    i := 0
    j := 0
    k := 0

    as := len(a)
    bs := len(b)

    r := make([]int, as+bs)

    for ; i < as && j < bs; {

        if a[i] <= b[j] {
            r[k] = a[i]
            i++
        } else {
            r[k] = b[j]
            j++
        }

        k++
    }

    for ; i < as; i++ {
        r[k] = a[i]
        k++
    }
    for ; j < bs; j++ {
        r[k] = b[j]
        k++
    }

    return r
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    return median(merge(nums1, nums2))
}
