package main

import "fmt"

type Test struct {
	res         int
	target      int
	mountainArr []int
}

func main() {
	tests := []Test{
		{2, 3, []int{1, 2, 3, 4, 5, 3, 1}},
		{-1, 3, []int{0, 1, 2, 4, 2, 1}},
		{5, 2, []int{0,1,3,4,3,2,1}},
	}

	for _, test := range tests {
		mountainArr := MountainArray{test.mountainArr, 0}
		res := findInMountainArray2(test.target, &mountainArr)
		if res != test.res {
			fmt.Println("Fail: actual=",res,", desired=",test.res)
		} else if mountainArr.callsToGet > 100 {
			fmt.Println("Fail: ", mountainArr.callsToGet, " calls to mountainArr.get()")
		} else {
			fmt.Println("Pass")
		}
	}
}
