package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var set map[byte]bool
	max := 1

	for ix := range s {
		if len(s)-ix+1 < max {
			break
		}
		set = map[byte]bool{s[ix]: true}

		for k := ix + 1; k < len(s); k++ {
			if _, contains := set[s[k]]; contains {
				if len(set) > max {
					max = len(set)
				}
				break
			}

			set[s[k]] = true
		}

		if len(set) > max {
			max = len(set)
		}
	}

	return max
}
