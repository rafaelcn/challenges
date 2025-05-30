func myAtoi(s string) int {
    
    digits := 0
    leading := true
    
    signal := int64(1)
    number := int64(0)

    lookup := map[byte]int64{
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
        //fmt.Println(string(s[i]), "number", number)

        if v, ok := lookup[s[i]]; ok && digits <= 10 {
            if v == 0 && leading {
                continue
            }
            digits++
            leading = false

            number = number * 10 + v
        } else {
            break
        }
    }

    if signal*number > math.MaxInt32 { return math.MaxInt32 }
    if signal*number < math.MinInt32 { return math.MinInt32 }

    return int(number*signal)
}
