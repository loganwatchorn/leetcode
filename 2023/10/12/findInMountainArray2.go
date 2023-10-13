package main

func findLeftmost(l int, r int, success func(i int) bool) int {
	var c int
	for r-l > 1 {
		c = (l + r) / 2
		if success(c) { r = c } else { l = c }
	}
	return r
}

func findInMountainArray2(target int, mountainArr *MountainArray) int {
	summit := findLeftmost(0, mountainArr.length()-1,
		func(i int) bool { return mountainArr.get(i) >= mountainArr.get(i+1) })

	idx := findLeftmost(0, summit,
		func(i int) bool { return target <= mountainArr.get(i) })
	if mountainArr.get(idx) == target { return idx }
	if mountainArr.get(0) == target { return 0 }

	idx = findLeftmost(summit, mountainArr.length()-1,
		func(i int) bool { return target >= mountainArr.get(i) })
	if mountainArr.get(idx) == target { return idx }

	return -1
}
