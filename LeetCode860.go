package main

func lemonadeChange(bills []int) bool {
	//贪心：优先用10来找零
	five := 0
	ten := 0

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		}

		if bills[i] == 10 {
			if five == 0 {
				return false
			} else {
				five--
				ten++
			}
		}

		if bills[i] == 20 {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}
