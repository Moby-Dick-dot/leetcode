package main

func moveZeroes(nums []int) {
	total := 0
	//统计零的个数
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			total++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
		}
	}

	for total != 0 {
		nums[len(nums)-total] = 0
		total--
	}
}
