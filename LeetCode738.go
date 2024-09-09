package main

import "strconv"

func monotoneIncreasingDigits(n int) int {
	//先将整数转成字符串
	str := strconv.Itoa(n)
	//然后将字符串转成字符数组，要不然不能修改
	ss := []byte(str)

	flag := len(ss)
	//然后从后往前！！！！遍历这个字符串(因为要i-1,所以要i>0)
	for i := len(ss) - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1]--
			flag = i
		}
	}

	//将flag之后的所有位置都改成9
	for i := flag; i < len(ss); i++ {
		ss[i] = '9'
	}

	str = string(ss)
	res, _ := strconv.Atoi(str)
	return res
}
