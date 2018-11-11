package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {

	for ix := 0; ix < len(nums)-1; ix++ {
		if nums[ix] != nums[ix+1] {
			continue
		}
		c := ix + 1
		for ; c < len(nums) && nums[ix] == nums[c]; c++ {
		}

		nums = append(nums[:ix+1], nums[c:]...)
	}
	return len(nums)
}
