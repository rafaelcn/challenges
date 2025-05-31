func repeatedCharacter(s string) byte {
    r := 0
    m := 0

    for i := range s {
        m = 1 << (s[i] - 'a')

        if r & m != 0 {
            return s[i]
        }

        r |= m
    }

    return byte(0)
}
