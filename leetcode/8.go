func myAtoi(s string) int {
    
    digits := []int{}
    signal := 1
    consumed := 0
    leading := true

    lookup := map[byte]int{
        '0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
    }

    // consume all blank spaces and the number signal first
    i := 0
    for i < len(s) && s[i] == ' ' { 
        i++
    }
    if i < len(s) && s[i] == '-' { 
        signal = -1 
        i++
    } else if i < len(s) && s[i] == '+' { 
        i++
    }

    // consume printable characters now
    for ; i < len(s); i++ {
        //fmt.Println(string(s[i]), digits, "i", i, "consumed", consumed)

        if v, ok := lookup[s[i]]; ok && len(digits) <= 10 {
            if v == 0 && leading {
                continue
            }
            consumed++
            leading = false
            digits = append(digits, v)
        } else {
            break
        }
    }

    if len(digits) == 0 { return 0 }
    
    number := int64(0)
    for i := 0; i < len(digits); i++ {
        e := 1 
        m := len(digits)-1-i

        for m > 0 {
            e *= 10
            m--
        }

        number += int64(digits[i]*e)
    }

    number = number*int64(signal)

    if number > math.MaxInt32 { number = math.MaxInt32 }
    if number < math.MinInt32 { number = math.MinInt32 }

    return int(number)
}
