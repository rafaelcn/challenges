type Transaction struct {
    Name string
    Time int
    Amount int
    City string

    Index int
}

func (t Transaction) String() string {
    return t.Name+","+strconv.Itoa(t.Time)+","+strconv.Itoa(t.Amount)+","+t.City
}

func (t Transaction) Difference(o Transaction) bool {
    r := t.Time - o.Time
    if r < 0 {
        return -r <= 60
    }

    return r <= 60
}

func invalidTransactions(transactions []string) []string {
    ts := map[string][]Transaction{}
    
    invalid := []string{}
    mask := make([]bool, len(transactions))

    for i, transaction := range transactions {
        properties := strings.Split(transaction, ",")

        name := properties[0]
        time, _ := strconv.Atoi(properties[1])
        amount, _ := strconv.Atoi(properties[2])
        city := properties[3]

        t := Transaction{Name: name, Time: time, Amount: amount, City: city, Index:i}
        
        if amount > 1000 {
            mask[i] = true
        }

        if _, ok := ts[name]; !ok {
            ts[name] = []Transaction{t}
        } else {
            ts[name] = append(ts[name], t)
        }
    }

    for _, tr := range ts {
        for i := 0; i < len(tr)-1; i++ {
            for j := i+1; j <= len(tr)-1; j++ {
                if tr[i].City == tr[j].City { continue }
           
                if tr[i].Difference(tr[j]) {
                    mask[tr[i].Index] = true
                    mask[tr[j].Index] = true
                }
            }
        }
    }

    for i, b := range mask {
        if b == true {
            invalid = append(invalid, transactions[i])
        }
    }

    return invalid 
}
