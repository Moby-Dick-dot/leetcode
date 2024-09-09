package main

//func partitionLabels(s string) []int {
//	//统计字符出现的最远距离（使用哈希表）
//	mp := make(map[uint8]int)
//	for i := 0; i < len(s); i++ {
//		mp[s[i]-'a'] = i
//	}
//
//	//使用左右边界来计算片段长度
//	result := make([]int, 0)
//	left, right := 0, 0
//
//	for i := 0; i < len(s); i++ {
//		right = max(right, mp[s[i]-'a'])
//		if i == right {
//			result = append(result, right-left+1)
//			left = i + 1
//		}
//	}
//
//	return result
//}
