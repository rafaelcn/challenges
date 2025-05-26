func addBinary(a string, b string) string {
    s := []byte{}

    i := len(a)-1
    j := len(b)-1
    carry := byte(0)

    for i >= 0 || j >= 0 || carry != 0 {
        if i >= 0 {
            carry += a[i] - '0' // 48-48 or 49-48
            i--
        }
        if j >= 0 {
            carry += b[j] - '0'
            j--
        }

        // carry = 0 -> 0 carry 0
        // carry = 1 -> 1 carry 0
        // carry = 2 -> 0 carry 1
        // carry = 3 -> 1 carry 1

        s = append([]byte{(carry%2)+48}, s...)
        carry /= 2
    }

    return string(s)
}
