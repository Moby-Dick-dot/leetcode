package main

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	cur := 0    //当前覆盖范围
	next := 0   //下一步覆盖范围
	result := 0 //更新覆盖范围的次数

	for i := 0; i < len(nums)-1; i++ {
		next = max(next, i+nums[i]) //保存最大的下一步覆盖范围
		if i == cur {               //i到达当前覆盖范围的最后
			cur = next              //更新当前覆盖范围
			result++                //记录更新范围的次数（跳跃次数）
			if cur >= len(nums)-1 { //更新完当前覆盖范围后要进行判断，看是否已经覆盖最后的元素
				break
			}
		}
	}

	return result
}
