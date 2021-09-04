package intersection

func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := make([][]int, 0)
	for _, pA := range A {
		for _, pB := range B {
			if pB[0] > pA[1] {
				break
			}
			d := diff(pA, pB)
			if len(d) > 0 {
				result = append(result, d)
			}
		}
	}
	return result
}

func diff(f, s []int) []int {
	fa, fb := f[0], f[1]
	sa, sb := s[0], s[1]

	if fa == sa && fb == sb {
		return f
	}
	if fa > sa && fb < sb {
		return f
	}
	if sa > fa && sb < fb {
		return s
	}
	if fa == sa {
		if fb < sb {
			return f
		}
		return s
	}
	if fb == sb {
		if fa > sa {
			return f
		}
		return s
	}
	if fb == sa {
		return []int{fb, sa}
	}
	if fa == sb {
		return []int{fa, sb}
	}
	if fa > sa && fa < sb {
		return []int{fa, sb}
	}
	if sa > fa && sa < fb {
		return []int{sa, fb}
	}
	return nil
}
