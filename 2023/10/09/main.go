package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"time"
)

const numTests = 100
const maxN = 100000
const minNum = -1000000000
const maxNum = 1000000000

type Test struct {
	N       int
	Correct bool
	Time    float64
	Space   int
}

func main() {
	tests := make([]Test, numTests)

	var N int
	var targetStart int
	var targetEnd int

	var nums []int
	var target int

	for testIdx := 0; testIdx < numTests; testIdx++ {
		N = maxN * testIdx / (numTests - 1)

		nums = make([]int, N)
		for i := 0; i < N; i++ {
			nums[i] = rand.Intn(maxNum-minNum+1) + minNum
		}
		sort.Ints(nums)

		switch testIdx {
		case 0, 20: // target not in nums
			targetStart, targetEnd = -1, -1
		case 21: // target is last
			targetStart, targetEnd = N-1, N-1
		case 22: // target is first
			targetStart, targetEnd = 0, 0
		default:
			if rand.Intn(2) == 0 {
				targetStart, targetEnd = rand.Intn(N), rand.Intn(N)
				if targetEnd < targetStart {
					targetStart, targetEnd = targetEnd, targetStart
				}
			} else {
				targetStart, targetEnd = -1, -1
			}
		}

		if targetStart == -1 || N == 0 {
			target = rand.Intn(maxNum-minNum+1) + minNum
		} else {
			target = nums[rand.Intn(targetEnd-targetStart+1)+targetStart]
			for i := targetStart; i <= targetEnd; i++ {
				nums[i] = target
			}
		}

		var timeStart time.Time
		var timeResult float64
		var spaceStart, spaceEnd runtime.MemStats
		var spaceResult int

		// record start time
		timeStart = time.Now()
		runtime.ReadMemStats(&spaceStart)

		result := searchRange(nums, target)

		timeResult = time.Since(timeStart).Seconds()
		runtime.ReadMemStats(&spaceEnd)
		spaceResult = int(spaceEnd.Alloc - spaceStart.Alloc)

		correct := result[0] == targetStart && result[1] == targetEnd
		if !correct {
			println("Test", testIdx, "failed")
			println("  Output: (", result[0], ",", result[1], ")")
			println("  Actual: (", targetStart, ",", targetEnd, ")")
			println("  nums: ", nums)
		}

		tests[testIdx] = Test{
			N:       N,
			Correct: correct,
			Time:    timeResult,
			Space:   spaceResult,
		}

		fmt.Printf("Test %d: %d, %t, %f, %d\n", testIdx, N, correct, timeResult, spaceResult)
	}

	numCorrect := 0
	for _, test := range tests {
		if test.Correct {
			numCorrect++
		}
	}

	fmt.Println("Correct:", numCorrect, "/", numTests)

	// Output results to CSV file results.csv
	file, _ := os.Create("results.csv")
	fmt.Fprintf(file, "N,Correct,Time (s),Space (b)\n")
	for _, test := range tests {
		fmt.Fprintf(file, "%d,%t,%f,%d\n", test.N, test.Correct, test.Time, test.Space)
	}
}
