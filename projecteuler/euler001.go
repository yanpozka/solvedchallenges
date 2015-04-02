package main

import (
	"fmt"
)

func main() {
	var T, N int64
	fmt.Scanf("%d", &T)

	for ; T > 0; T-- {
		fmt.Scanf("%d", &N)

		m_3 := N / 3
		m_5 := N / 5
		m_15 := N / 15

		if N%3 == 0 {
			m_3--
		}
		if N%5 == 0 {
			m_5--
		}
		if N%15 == 0 {
			m_15--
		}
		// fmt.Println(m_3, m_5, m_15)
		fmt.Println(sm(m_3)*3 + sm(m_5)*5 - sm(m_15)*15)
	}
}

func sm(n int64) int64 {
	return (n * (n + 1)) / 2
}
