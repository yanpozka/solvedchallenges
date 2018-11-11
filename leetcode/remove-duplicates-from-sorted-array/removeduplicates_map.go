package removeduplicatesfromsortedarray

func removeDuplicatesMap(nums []int) int {
	m := make(map[int]bool)
	var count int
	for _, n := range nums {
		if !m[n] {
			nums[count] = n
			count++
		}
		m[n] = true
	}
	nums = nums[:len(m)]
	return len(m)
}
