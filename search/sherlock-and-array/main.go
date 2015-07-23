package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0 && scanner.Scan(); T-- {
		N, _ := strconv.Atoi(scanner.Text())

		ac_left := make([]uint64, N)
		ac_right := make([]uint64, N)
		ns := make([]uint64, N)

		for ix := 0; ix < N && scanner.Scan(); ix++ {
			nn, _ := strconv.Atoi(scanner.Text())
			var num uint64 = uint64(nn)
			ac_left[ix] = num
			if ix > 0 {
				ac_left[ix] += ac_left[ix-1]
			}
			ns[ix] = num
		}

		ac_right[N-1] = ns[N-1]

		for ix := N - 2; ix >= 0; ix-- {
			ac_right[ix] = ns[ix] + ac_right[ix+1]
		}

		var ok bool
		if N > 1 {
			if ac_right[1] == 0 {
				fmt.Println("YES")
				ok = true
			}
		} else {
			fmt.Println("YES")
			ok = true
		}
		if ok {
			continue
		}
		for ix := 1; ix < N-1; ix++ {
			if ac_left[ix-1] == ac_right[ix+1] {
				fmt.Println("YES")
				ok = true
				break
			}
		}
		if ok {
			continue
		}
		if N-2 >= 0 {
			if ac_left[N-2] == 0 {
				fmt.Println("YES")
				ok = true
			}
		} else {
			fmt.Println("YES")
			ok = true
		}
		if !ok {
			fmt.Println("NO")
		}
	} // Tests
}
