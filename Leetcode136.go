package main

func singleNumber(nums []int) int {
	mp := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		mp[nums[i]] += 1
	}

	for key, value := range mp {
		if value == 1 {
			return key
		}
	}

	return 0
}
