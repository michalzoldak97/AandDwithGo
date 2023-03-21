package minmax

import (
	"runtime"
)

func FindNaive(nums []int) (int, int) {
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if num < min {
			min = num
			continue
		}

		if num > max {
			max = num
		}
	}

	return min, max
}

func findMax(nums []int, res chan int) {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	res <- max
}

func findMin(nums []int, res chan int) {
	min := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	res <- min
}

func FindNaiveDouble(nums []int) (int, int) {
	minMaxChan := make(chan int, 2)

	go findMin(nums, minMaxChan)
	go findMax(nums, minMaxChan)

	res1, res2 := <-minMaxChan, <-minMaxChan

	if res1 < res2 {
		return res1, res2
	}

	return res2, res1
}

func FindOptim(nums []int) (int, int) {
	min, max, i := 0, 0, 0

	l := len(nums)

	if l%2 != 0 {
		min = nums[0]
		max = nums[0]
		i = 1
	} else {
		if nums[0] < nums[1] {
			min = nums[0]
			max = nums[1]
		} else {
			min = nums[1]
			max = nums[0]
		}
		i = 2
	}

	for i < l {
		if nums[i] < nums[i+1] {
			if nums[i] < min {
				min = nums[i]
			}

			if nums[i+1] > max {
				max = nums[i+1]
			}
		} else {
			if nums[i] > max {
				max = nums[i]
			}
			if nums[i+1] < min {
				min = nums[i+1]
			}
		}

		i += 2
	}

	return min, max
}

func FindNaiveCPU(nums []int) (int, int) {
	nSize := len(nums)
	cores := runtime.NumCPU()

	if nSize < cores {
		return FindNaive(nums)
	}

	step := nSize / (cores / 2)
	mod := nSize % cores

	minMaxChan := make(chan int, cores)
	defer close(minMaxChan)

	results := make([]int, (cores + mod))

	for i := 0; i < (cores / 2); i++ {
		go findMin(nums[(i*step):((step+i*step)-1)], minMaxChan)
		go findMax(nums[(i*step):((step+i*step)-1)], minMaxChan)
	}

	for i := 0; i < cores; i++ {
		results[i] = <-minMaxChan
	}

	if mod == 0 {
		return FindNaive(results)
	}

	mIdx := 1
	for i := cores; i < (cores + mod); i++ {
		results[i] = nums[(nSize - mIdx)]
		mIdx++
	}

	return FindNaive(results)
}
