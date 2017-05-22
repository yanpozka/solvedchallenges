package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scanf("%d%d%d", &n, &m, &k)

	mr := make(map[int][]*tuple)
	for ; k > 0; k-- {
		var r, c1, c2 int
		fmt.Scanf("%d%d%d", &r, &c1, &c2)

		if _, contains := mr[r]; !contains {
			mr[r] = []*tuple{&tuple{Min: c1, Max: c2}}
			continue
		}

		var ok bool
		for _, p := range mr[r] {
			if (c1 >= p.Min && c1 <= p.Max) || (c2 >= p.Min && c2 <= p.Max) {
				if p.Min > c1 {
					p.Min = c1
				}
				if p.Max < c2 {
					p.Max = c2
				}
				ok = true
			}
		}
		if ok == false {
			mr[r] = append(mr[r], &tuple{Min: c1, Max: c2})
		}
	}

	var result int64
	for _, tuples := range mr {
		var pr int
		for _, p := range tuples {
			pr += p.Max - p.Min + 1
		}
		result += int64(m - pr)
	}

	fmt.Println(result + int64((n-len(mr))*m))
}

type tuple struct {
	Min, Max int
}
