package product

func productExceptSelf(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}

	prodsFirst := make([]int, len(nums))
	prodsFirst[0] = nums[0]
	for ix := 1; ix < len(nums); ix++ {
		prodsFirst[ix] = prodsFirst[ix-1] * nums[ix]
	}

	prodsLast := make([]int, len(nums))
	prodsLast[len(nums)-1] = nums[len(nums)-1]
	for ix := len(nums) - 2; ix >= 0; ix-- {
		prodsLast[ix] = prodsLast[ix+1] * nums[ix]
	}

	sol := make([]int, len(nums))
	sol[0] = prodsLast[1]
	sol[len(nums)-1] = prodsFirst[len(prodsFirst)-2]

	for ix, lim := 1, len(sol)-1; ix < lim; ix++ {
		sol[ix] = prodsFirst[ix-1] * prodsLast[ix+1]
	}

	return sol
}
