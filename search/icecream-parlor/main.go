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
	T, _ := strconv.Atoi(scanner.Text())
	for ; T > 0; T-- {
		scanner.Scan()
		M, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())

		nums := make(IcecreamSlice, 0, N)

		for ix := 0; ix < N && scanner.Scan(); ix++ {
			x, _ := strconv.Atoi(scanner.Text())
			if x < M {
				nums = append(nums, &ice{Val: x, Pos: ix + 1})
			}
		}
		sort.Sort(nums) // !!
		N = len(nums)

		for i, x := range nums {
			ln, val := N-(i+1), M-x.Val
			tmp := nums[i+1:]

			if x.Val >= M/2 {
				tmp = nums[:i]
				ln = i
			}

			if pos := sort.Search(ln, func(idx int) bool { return tmp[idx].Val >= val }); pos != ln {
				if x.Val+tmp[pos].Val == M {
					if x.Pos < tmp[pos].Pos {
						fmt.Println(x.Pos, tmp[pos].Pos)
					} else {
						fmt.Println(tmp[pos].Pos, x.Pos)
					}
					break
				}
			}
		}
	}
}

type ice struct {
	Pos, Val int
}

func (ic *ice) String() string {
	return fmt.Sprintf("(Pos:%d Val:%d)", ic.Pos, ic.Val)
}

type IcecreamSlice []*ice

func (p IcecreamSlice) Len() int           { return len(p) }
func (p IcecreamSlice) Less(i, j int) bool { return p[i].Val < p[j].Val }
func (p IcecreamSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
