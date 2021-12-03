func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if v >= 0 || k == 0 {
			break
		}
		nums[i] = -nums[i]
		k--
	}
	sort.Ints(nums)
	k = k % 2
	ans := 0
	for i, num := range nums {
		if i == 0 && k == 1 {
			num = -num
		}
		ans += num
	}
	return ans
}

