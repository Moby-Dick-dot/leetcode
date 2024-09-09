package main

func candy(ratings []int) int {
	//两边的情况：要先比较一边再比较另一边
	candys := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		candys[i] = 1
	}

	//右孩比左孩得分高的情况(从前往后遍历)
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		}
	}

	//左孩比右孩的分高的情况（从后往前遍历）
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i+1]+1, candys[i])
		}
	}

	res := 0
	for i := 0; i < len(ratings); i++ {
		res += candys[i]
	}

	return res
}
