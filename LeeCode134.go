package main

func canCompleteCircuit(gas []int, cost []int) int {
	//计算剩余油量，然后进行求和，若和为负数，则从0继续开始

	curSum := 0
	totalSum := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]

		if curSum < 0 {
			curSum = 0
			start = i + 1
		}
	}

	if totalSum < 0 {
		return -1
	}

	return start
}
