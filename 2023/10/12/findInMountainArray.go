package main

func findInMountainArray(target int, mountainArr *MountainArray) int {
	var l, c, r int

	l, r = 0, mountainArr.length()-1
	for l+1 < r {
		c = (l + r) / 2
		if mountainArr.get(c) < mountainArr.get(c+1) {
			l = c
		} else {
			r = c
		}
	}
	summit := r

	l, r = 0, summit
	for l+1 < r {
		c = (l + r) / 2
		if mountainArr.get(c) < target {
			l = c
		} else {
			r = c
		}
	}
	if mountainArr.get(r) == target {
		return r
	}
	if mountainArr.get(0) == target {
		return 0
	}

	l, r = summit, mountainArr.length()-1
	for l+1 < r {
		c = (l + r) / 2
		if mountainArr.get(c) > target {
			l = c
		} else {
			r = c
		}
	}
	if mountainArr.get(r) == target {
		return r
	}

	return -1
}
