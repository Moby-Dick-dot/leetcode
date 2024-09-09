package main

//import "sort"
//
//func reconstructQueue(people [][]int) [][]int {
//	sort.Slice(people, func(i, j int) bool {
//		if people[i][0] == people[j][0] {
//			return people[i][1] < people[j][1]
//		}
//		return people[i][0] > people[j][0]
//	})
//
//	//创建一个新的队列
//	queue := make([][]int, 0, len(people))
//
//	//然后根据根据序号插入到队列当中
//	for _, p := range people {
//		index := p[1]
//		queue = append(queue[:index], append([][]int{p}, queue[index:]...)...)
//	}
//	return queue
//}
