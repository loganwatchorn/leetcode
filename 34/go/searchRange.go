package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l := 0
	r := len(nums) - 1
	var c int

	res := make([]int, 2, 2)

	if nums[l] == target {
		res[0] = 0
	} else {
		for l+1 < r {
			c = (l + r) / 2
			if nums[c] < target {
				l = c
			} else {
				r = c
			}
		}

		if nums[r] != target {
			return []int{-1, -1}
		}

		res[0] = r
	}

	r = len(nums) - 1

	if target == nums[r] {
		res[1] = r
	} else {
		for l+1 < r {
			c = (l + r) / 2
			if nums[c] > target {
				r = c
			} else {
				l = c
			}
		}

		res[1] = l
	}

	return res
}
