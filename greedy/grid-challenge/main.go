package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	for T, _ := strconv.Atoi(scanner.Text()); T > 0 && scanner.Scan(); T-- {
		N, _ := strconv.Atoi(scanner.Text())

		matrix := make([][]byte, N)

		for ix := 0; ix < N && scanner.Scan(); ix++ {
			bs := []byte(scanner.Text())
			sort.Sort(ByteSlice(bs))
			matrix[ix] = bs
		}

		var ok bool = true
		for c := 0; c < N && ok; c++ {
			bc := make([]byte, N)

			for k := 0; k < N; k++ {
				bc[k] = matrix[k][c]
			}

			if is_sorted := sort.IsSorted(ByteSlice(bc)); is_sorted == false {
				ok = is_sorted
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

//
type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
