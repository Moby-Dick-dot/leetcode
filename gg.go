package main

//func getRightBorder(nums []int, target int) int {
//	left := 0
//	right := len(nums) - 1
//	RightBorder := -2
//	for left <= right {
//		medium := left + (right-left)/2
//		if nums[medium] > target {
//			right = medium - 1
//		} else {
//			left = medium + 1
//			RightBorder = left
//		}
//	}
//	return RightBorder
//}
//
//func getLeftBorder(nums []int, target int) int {
//	left := 0
//	right := len(nums) - 1
//	LeftBorder := -2
//	for left <= right {
//		medium := left + (right-left)/2
//		if nums[medium] >= target {
//			right = medium - 1
//			LeftBorder = right
//		} else {
//			left = medium + 1
//		}
//	}
//	return LeftBorder
//}
//
//func searchRange(nums []int, target int) []int {
//	left := getLeftBorder(nums, target)
//	right := getRightBorder(nums, target)
//	if left == -2 || right == -2 {
//		return []int{-1, -1}
//	}
//	if right-left > 1 {
//		return []int{left + 1, right - 1}
//	}
//	return []int{-1, -1}
//}

//func mySqrt(x int) int {
//	left := 0
//	right := x
//	for left <= right {
//		medium := left + (right-left)/2
//		if medium*medium == x {
//			return medium
//		} else if medium*medium > x && (medium-1)*(medium-1) < x {
//			return medium - 1
//		} else if medium*medium > x {
//			right = medium - 1
//		} else if medium*medium < x {
//			left = medium + 1
//		}
//	}
//	return -1
//}

//func isPerfectSquare(num int) bool {
//	left := 0
//	right := num
//	for left <= right {
//		medium := left + (right-left)/2
//		if medium*medium == num {
//			return true
//		} else if medium*medium > num {
//			right = medium - 1
//		} else if medium*medium < num {
//			left = medium + 1
//		}
//	}
//	return false
//}

//func removeElement(nums []int, val int) int {
//	slow, fast := 0, 0
//	for fast = 0; fast < len(nums); fast++ {
//		if nums[fast] != val {
//			nums[slow] = nums[fast]
//		}
//		slow++
//	}
//	return slow
//}

//func searchInsert(nums []int, target int) int {
//	left := 0
//	right := len(nums) - 1
//
//	if left <= right {
//		medium := left + (right-left)/2
//		if nums[medium] > target {
//			right = medium - 1
//		} else if nums[medium] < target {
//			left = medium + 1
//		} else {
//			return medium
//		}
//	}
//	return right + 1
//}

//func removeElement(nums []int, val int) int {
//	fast, slow := 0, 0
//	for fast = 0; fast < len(nums); fast++ {
//		if nums[fast] != val {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//	return slow
//}

//func sortedSquares(nums []int) []int {
//	result := []int{}
//	k := len(nums) - 1
//	for i, j := 0, len(nums)-1; i <= j; {
//		if nums[i]*nums[i] > nums[j]*nums[j] {
//			result[k] = nums[i] * nums[i]
//			k--
//			i++
//		} else {
//			result[k] = nums[j] * nums[j]
//			k--
//			j--
//		}
//	}
//	return result
//
//}
//
//func minSubArrayLen(target int, nums []int) int {
//	sum := 0
//	subL := 0
//	i := 0
//	result := math.MaxInt
//	for j := 0; j < len(nums); j++ {
//		sum = sum + nums[j]
//		for sum >= target {
//			subL = j - i + 1
//			result = min(result, subL)
//			sum = sum - nums[i]
//			i++
//		}
//	}
//	if result == math.MaxInt {
//		return 0
//	}
//	return result
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func removeElements(head *ListNode, val int) *ListNode {
//	for head.Val == val {
//		head = head.Next
//	}
//	cur := head
//	for cur.Next != nil {
//		if cur.Next.Val == val {
//			cur.Next = cur.Next.Next
//		} else {
//			cur = cur.Next
//		}
//	}
//	return head
//}

//	type MyLinkedList struct {
//		Val  int
//		Next *ListNode
//	}
//
//	func Constructor() MyLinkedList {
//		return MyLinkedList{0, &MyLinkedList{}}
//	}
//
//	func (this *MyLinkedList) Get(index int) int {
//		if index < 0 || index > this.size {
//			return -1
//		}
//		dummyhead := new(MyLinkedList) //创建虚拟头节点
//		dummyhead.Next = this
//		cur := this
//		for index != 0 {
//			cur = cur.Next
//			index--s
//		}
//		return cur.Val
//	}
//
//	func (this *MyLinkedList) AddAtHead(val int) {
//		dummyhead := new(MyLinkedList)
//		dummyhead.Next = this
//		newNode := new(MyLinkedList)
//		newNode.Next = dummyhead.Next
//		dummyhead.Next = newNode
//	}
//
//	func (this *MyLinkedList) AddAtTail(val int) {
//		dummyhead := new(MyLinkedList)
//		dummyhead.Next = this
//		newNode := new(MyLinkedList)
//		cur := dummyhead
//		for cur.Next != nil {
//			cur = cur.Next
//		}
//		cur.Next = newNode
//	}
//
//	func (this *MyLinkedList) AddAtIndex(index int, val int) {
//		dummyhead := new(MyLinkedList)
//		dummyhead.Next = this
//		newNode := new(MyLinkedList)
//		cur := dummyhead //很重要！！！！ 第n个节点一定要是cur.next
//		for index != 0 {
//			cur = cur.Next
//			index--
//		}
//		newNode.Next = cur.Next
//		cur.Next = newNode
//	}
//
//	func (this *MyLinkedList) DeleteAtIndex(index int) {
//		dummyhead := new(MyLinkedList)
//		dummyhead.Next = this
//		cur := dummyhead //不能=dead，这样可以保证找到第n个节点的前一个节点
//		for index != 0 {
//			cur = cur.Next
//		}
//		cur.Next = cur.Next.Next
//	}
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//func reverseList(head *ListNode) *ListNode {
//	pre := nil
//	cur := head
//	for cur != nil {
//		temp := cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = temp
//	}
//	return pre
//}

//func reverseList(head *ListNode) *ListNode {
//
//	return reverse(nil, head)
//
//}
//
//func reverse(pre *ListNode, cur *ListNode) *ListNode {
//	if cur == nil {
//		return pre
//	}
//	temp := cur.Next
//	cur.Next = pre
//	return reverse(cur, temp)
//}

//	type ListNode struct {
//		Val  int
//		Next *ListNode
//	}
//
//	func swapPairs(head *ListNode) *ListNode {
//		dummyhead := new(ListNode)
//		dummyhead.Next = head
//		cur := dummyhead
//
//		for cur.Next != nil && cur.Next.Next != nil {
//			temp := cur.Next            //存节点1的指针
//			temp1 := cur.Next.Next.Next //提前存节点3的指针，要不然找不到了
//			cur.Next = cur.Next.Next
//			cur.Next.Next = temp
//			temp.Next = temp1
//			cur = cur.Next.Next //移动cur，指向下一组交换节点的前一个节点
//
//		}
//		return dummyhead.Next
//	}
//

//
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func detectCycle(head *ListNode) *ListNode {
//	fast := head
//	slow := head
//	for fast != nil && fast.Next != nil {
//		fast = fast.Next.Next
//		slow = slow.Next
//		if fast == slow {
//			index1 := fast
//			index2 := head
//			for index1 != index2 {
//				index1 = index1.Next
//				index2 = index2.Next
//			}
//			return index2
//		}
//	}
//	return nil
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//func detectCycle(head *ListNode) *ListNode {
//	fast := head
//	slow := head
//	for fast != nil && fast.Next != nil{
//		fast = fast.Next.Next //快指针走两步
//		slow = slow.Next//慢指针走一步
//		if slow == fast{ //快慢指针相遇
//			index1 := fast
//			index2 := head
//			for index1 != index2{ //找到入口位置
//				index1 = index1.Next
//				index2 = index2.Next
//			}
//			return index1
//		}
//	}
//	return nil
//}

//func groupAnagrams(strs []string) [][]string {
//	mp := map[[26]int][]string{}
//	for _, str := range strs {
//		cnt := [26]int{}
//		for _, ch := range str {
//			cnt[ch-'a']++
//		}
//		mp[cnt] = append(mp[cnt], str)
//	}
//
//	ans := make([][]string, 0, len(mp))
//	for _, v := range mp {
//		ans = append(ans, v)
//	}
//	return ans
//}

//	func intersection(nums1 []int, nums2 []int) []int {
//		mp := map[int]int{}
//		for i := 0; i < len(nums1); i++ {
//			mp[nums1[i]]++
//		}
//		res := []int{}
//		for i := 0; i < len(nums2); i++ {
//			if mp[nums2[i]] != 0 {
//				res = append(res, nums2[i])
//				mp[nums2[i]] = 0
//			}
//		}
//		return res
//	}

//func isHappy(n int) bool {
//	m := make(map[int]bool) //创建一个空的哈希表
//
//	//通过哈希表判断是否存在循环，如果哈希表中已经存在该键说明出现循环
//	for n != 1 && !m[n] { //如果n不为1且键n不存在于哈希表中
//		m[n] = true   //如果哈希表中不存在键n，那么将其置为存在
//		n = getSum(n) //更新n
//	}
//	return n == 1 //两种情况，第一种n = 1，第二种n 不等于1，是因为出现循环，哈希表已经存在键n
//}
//
//func getSum(n int) int {
//	sum := 0
//	//整个算法可以依次取到每个位置的数字
//	for n != 0 {
//		sum += (n % 10) * (n % 10) //取个位数
//		n = n / 10                 //删掉个位数，只剩下前面几位
//	}
//	return sum
//}

//func twoSum(nums []int, target int) []int {
//	mp := make(map[int]int)
//	for index, val := range nums {
//		if preindex, ok := mp[target-val]; ok { //在map中寻找与当前元素相加等于target的元素
//			return []int{preindex, index} //如果存在返回二个元素的下标
//		} else {
//			mp[val] = index //如果不存在，将当前元素及其下标加入到map中
//		}
//	}
//	return []int{}
//}

//	func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
//		mp := make(map[int]int)
//
//		for _, val1 := range nums1 {
//			for _, val2 := range nums2 {
//				mp[val1+val2]++ //统计nums1和nums2元素之和及其对应出现的次数
//			}
//		}
//		res := 0
//		for _, val3 := range nums3 {
//			for _, val4 := range nums4 {
//				res += mp[0-val3-val4] //查哈希表 因为val1 + val2 + val3 + val4=0，所以找到0 - val3 - val4出现的次数，就相当于找到四数之和为0出现的次数
//			}
//		}
//		return res
//	}
//func canConstruct(ransomNote string, magazine string) bool {
//	hash := [26]int{}
//	for _, ch := range magazine {
//		hash[ch-'a']++ //构建一张哈希表，统计magazine中字母出现的次数
//	}
//
//	for _, rans := range ransomNote { //检查ransomNote中的每个字母是否都存在于magazine中
//		hash[rans-'a']--        //查看magazine中是否有字母可以组成ransomNote
//		if hash[rans-'a'] < 0 { //如果maganize中缺了字母，即不能组成
//			return false
//		}
//	}
//	return true
//}

//func threeSum(nums []int) [][]int {
//	//定义一个二维切片，用来存放结果三元组
//	res := [][]int{}
//	//首先要对数组进行排序
//	sort.Ints(nums)
//	//去重
//	i := 0
//	for i > 0 && nums[i] == nums[i+1] {
//		i++
//	}
//	//然后对排序后的数组进行遍历
//	for ; i < len(nums); i++ {
//		left := i + 1          //left指向i+1位置的值
//		right := len(nums) - 1 //right指向数组的最后一个元素
//		sum := nums[i] + nums[left] + nums[right]
//		for left != right { //直到左右指针相遇
//			for left < right && nums[left] == nums[left+1] {
//				left++
//			}
//			for left < right && nums[right] == nums[right-1] {
//				right--
//			}
//			if sum == 0 {
//				res = append(res, []int{nums[i], nums[left], nums[right]})
//				left++
//				right--
//			} else if sum > 0 {
//				right--
//			} else if sum < 0 {
//				left++
//			}
//			//if sum > 0 {
//			//	right-- //和太大，右指针左移
//			//} else if sum < 0 {
//			//	left++ //和太小，左指针右移
//			//} else if sum == 0 {
//			//	res = append(res, []int{nums[i], nums[left], nums[right]}) //和为0，将该三元组加入结果二维切片中
//			//}
//		}
//	}
//	return res
//}
//func fourSum(nums []int, target int) [][]int {
//	res := [][]int{}
//
//	for i := 0; i < len(nums - 3); i++ {
//		for j := 1; j < len(nums - 2); j++ {
//			if 	j
//		}
//	}
//
//}

//	func numberReplace(s string) string {
//		sb := []byte(s) //将字符串转成切片
//		oldsize := len(sb)
//		//首先，找到字符串中有多少个数字
//		count := 0
//		for i := 0; i < len(sb); i++ {
//			if sb[i] >= '0' && sb[i] <= '9' {
//				count++
//			}
//		}
//		//增加切片长度
//		for count > 0 {
//			sb = append(sb, []byte("     ")...)
//			count--
//		}
//		//双指针：左右指针，从后往前进行更新
//		left := oldsize - 1
//		right := len(sb) - 1
//		tmpstrs := []byte("number")
//		for left < right {
//			rightshift := 1
//			//如果是数字则加入number
//			if sb[left] >= '0' && sb[left] <= '9' {
//				for i, tmpstr := range tmpstrs {
//					sb[right-len(tmpstrs)+1+i] = tmpstr
//				}
//				rightshift = len(tmpstrs)
//			} else {
//				sb[right] = sb[left]
//			}
//			//更新指针
//			right -= rightshift
//			left -= 1
//		}
//
//		return string(sb)
//	}
//
//	func reverseString(s []byte) {
//		//双指针：左右指针
//		left := 0
//		right := len(s) - 1
//
//		for left < right {
//			s[left], s[right] = s[right], s[left]
//			//更新指针
//			right--
//			left++
//		}
//	}
//
//	func reverseStr(s string, k int) string {
//		sb := []byte(s)
//		length := len(s)
//		//找到要反转k字符的开头位置i，然后翻转切片
//		for i := 0; i < length; i += 2 * k {
//			if length-i >= k {
//				reverseString(sb[i : i+k])
//			} else {
//				reverseString(sb[i:length])
//			}
//		}
//		return string(sb)
//	}
//func repeatedSubstringPattern(s string) bool {
//	flag := false
//	for i := 1; i < len(s)/2; i++ {
//		if reflect.DeepEqual(s[:i], s[i+1:2*i]) {
//			flag = true
//		}
//	}
//	return flag
//}

//func reverseString(s []byte)  {
//	left := 0
//	right := len(s) - 1
//	for left < right {
//		s[left], s[right] = s[right], s[left]
//		left++
//		right--
//	}
//
//}
//
//// 右旋字符串
//func rightRote(s string, k int) string {
//
//	ss := []byte(s)
//	reverseString(ss)
//	reverseString(ss[:k])
//	reverseString(ss[k+1:])
//	return string(ss)
//}

// KMP算法：
// 第一步求出next数组
// 第二步根据next数组来减少不必要的比较
func getNext(next []int, s string) {
	//初始化：j指向前缀字符串末尾 ，i指向后缀字符串末尾（j还代表最长相等前后缀长度）
	j := 0
	next[0] = 0

	for i := 1; i < len(s); i++ {
		//前缀和后缀不等
		for j > 0 && s[j] != s[i] {
			j = next[j-1] //回退
		}
		//前缀和后缀相等
		if s[j] == s[i] {
			j++
		}
		//更新next数组
		next[i] = j
	}
}

func strStr(haystack string, needle string) int {
	n := len(needle)
	//得到next数组
	next := make([]int, n)
	getNext(next, needle)

	j := 0
	//根据数组来减少不必要的比较
	for i := 0; i < len(haystack); i++ {
		//如果遇到不相等的元素，回退
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] //回退
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

// 重复的子串
//import "strings"
//
//func repeatedSubstringPattern(s string) bool {
//	//首先在s后面再添加一个s
//	a := []byte(s)
//	ss := append(a, a...)
//	slow := 0
//	//然后去掉ss中的首尾字母
//	for i := 0; i < len(ss); i++ {
//		if i != 0 && i != len(ss)-1 {
//			ss[slow] = ss[i]
//			slow++
//		}
//	}
//	ss = ss[:slow]
//	//在ss中寻找是否出现过s
//	res := strings.Index(string(ss), s)
//	if res >= 0 {
//		return true
//	} else {
//		return false
//	}
//}

//// 链表相交
//// 分别统计两个链表长度，然后将长链表指针移动到和短链表对齐的位置
//// 然后依次对比两指针是否相同，直到为空
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	lengtha, lengthb := 0, 0
//	curA, curB := headA, headB
//	for curA != nil {
//		curA = curA.Next
//		lengtha++
//	}
//	for curB != nil {
//		curB = curB.Next
//		lengthb++
//	}
//	curA, curB = headA, headB
//	sub := math.Abs(float64(lengtha - lengthb))
//	//if lengtha > lengthb {
//	//	for sub > 0 {
//	//		curA = curA.Next
//	//		sub--
//	//	}
//	//} else {
//	//	for sub > 0 {
//	//		curB = curB.Next
//	//		sub--
//	//	}
//	//}
//	for sub > 0 {
//		if lengtha > lengthb{
//			curA = curA.Next
//		}else{
//			curB = curB.Next
//		}
//		sub--
//	}
//	for curA != curB && curA != nil {
//		curA = curA.Next
//		curB = curB.Next
//	}
//	return curA
//}

//// 三数之和
//// 哈希表不能用因为去重逻辑太复杂
//// 所以使用双指针来降低时间复杂度
//// 因为是三数，所以外面套一层循环，里面使用双指针
//func threeSum(nums []int) [][]int {
//	//对数组排序
//	sort.Ints(nums)
//
//	res := [][]int{}
//	//最外面一层循环
//	for i := 0; i < len(nums)-2; i++ {
//		//剪枝
//		if nums[i] > 0 {
//			return [][]int{}
//		}
//		//nums[i]去重
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		//双指针遍历
//		left := i + 1
//		right := len(nums) - 1
//		sum := nums[i] + nums[left] + nums[right]
//		for left < right {
//			if sum < 0 {
//				left++
//			} else if sum > 0 {
//				right--
//			} else {
//				res = append(res, []int{nums[i], nums[left], nums[right]})
//				//nums[left],nums[right]去重
//				for left < right && nums[left] == nums[left+1] {
//					left++
//				}
//				for left < right && nums[right] == nums[right-1] {
//					right--
//				}
//				left++
//				right--
//			}
//		}
//	}
//	return res
//}

// 四数之和
// 也使用双指针
// 和三数之和类似 只不过外面多套了一层循环
// 需要注意的是 剪枝的条件
//func fourSum(nums []int, target int) [][]int {
//	//对数组排序
//	sort.Ints(nums)
//
//	res := [][]int{}
//
//	//开始循环
//	for k := 0; k < len(nums)-3; k++ {
//
//		//去重
//		if k > 0 && nums[k] == nums[k-1] {
//			continue
//		}
//		for i := k + 1; i < len(nums)-2; i++ {
//
//			//去重
//			if i > k+1 && nums[i] == nums[i-1] {
//				continue
//			}
//			left := i + 1
//			right := len(nums) - 1
//
//			for left < right {
//				sum := nums[k] + nums[i] + nums[left] + nums[right]
//				if sum > target {
//					right--
//				} else if sum < target {
//					left++
//				} else {
//					res = append(res, []int{nums[k], nums[i], nums[left], nums[right]})
//					for left < right && nums[left] == nums[left+1] {
//						left++
//					}
//					for left < right && nums[right] == nums[right-1] {
//						right--
//					}
//					right--
//					left++
//
//				}
//			}
//		}
//	}
//	return res
////}
//
//func search(nums []int, target int) int {
//	//
//	left := 0
//	right := len(nums) - 1
//	for left <= right {
//		mid := left + (right-left)/2
//		if nums[mid] > target {
//			right = mid - 1
//		} else if nums[mid] < target {
//			left = mid + 1
//		} else {
//			return mid
//		}
//	}
//	return -1
//}

//func removeElement(nums []int, val int) int {
//	//保留符合条件的元素
//	slow := 0 //用来收集符合条件的元素
//	//遍历数组
//	for i := 0; i < len(nums); i++ {
//		if nums[i] != val {
//			nums[slow] = nums[i]
//			slow++
//		}
//	}
//	return slow
//}

//func sortedSquares(nums []int) []int {
//	//思路：不会
//	//双指针法
//	left := 0
//	right := len(nums) - 1
//	res := []int{}
//
//	for left <= right {
//		if nums[left]*nums[left] > nums[right]*nums[right] {
//			res = append(res, nums[left]*nums[left])
//			left++
//		} else {
//			res = append(res, nums[right]*nums[right])
//			right--
//		}
//	}
//	sort.Ints(res)
//	return res
//}
//
//func sortedSquares(nums []int) []int {
//	//创建一个和原始切片一样大的切片，从后往前进行赋值
//	res := make([]int, len(nums), len(nums))
//	k := len(nums) - 1
//
//}

//func minSubArrayLen(target int, nums []int) int {
//	//看到连续子数组，想到滑动窗口
//	//定义一个最小值，在滑动的过程中不断更新(通过比较)
//	res := math.MaxInt
//	subL := 0
//	//先扩大滑动窗口大小，然后再缩小
//	i := 0
//	j := 0
//	sum := 0
//	for j = 0; j < len(nums); j++ {
//		sum += nums[j]
//		for sum >= target {
//			subL = j - i + 1
//			res = min(res, subL)
//			sum -= nums[i]
//			i++
//		}
//	}
//	if res == math.MaxInt {
//		return 0
//	}
//	return res
//
//}
//
//func generateMatrix(n int) [][]int {
//	//定义一个n*n的切片
//	nums := make([][]int, n)
//	for i := 0; i < n; i++ {
//		nums[i] = make([]int, n)
//	}
//	//定义起始位置和偏移量
//	startx, starty := 0, 0
//	offset := 1
//	k := n / 2
//	i, j := 0, 0
//	count := 1
//	//循环多少圈
//	for k > 0 {
//		//每次循环一圈要处理4条边
//		//上边
//		//（i，j）
//		for j = starty; j < n-offset; j++ {
//			nums[startx][j] = count
//			count++
//		}
//		//上面循环结束，j已经变成n-offset（固定了）
//		//右边
//		for i = startx; i < n-offset; i++ {
//			nums[i][j] = count
//			count++
//		}
//		//上面循环结束，i已经变成n-offset(固定了)
//		//下边
//		for ; j > starty; j-- {
//			nums[i][j] = count
//			count++
//		}
//		//上面循环结束，j已经变成starty(固定了)
//		//左边
//		for ; i > startx; i-- {
//			nums[i][j] = count
//			count++
//		}
//		k--
//		startx++
//		starty++ //移动起始位置
//		offset++ //改变偏移量
//	}
//	if n%2 == 1 {
//		nums[i][j] = count
//	}
//	return nums
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 要保存它的前一个节点
//func removeElements(head *ListNode, val int) *ListNode {
//	dummyhead := new(ListNode)
//
//	pre := dummyhead
//	cur := head
//
//	for cur != nil {
//		if cur.Val == val {
//			pre.Next = cur.Next
//			cur = cur.Next
//		} else {
//			pre = cur
//			cur = cur.Next
//
//		}
//
//	}
//	return dummyhead.Next
//}

// 两两交换链表中的节点

//什么时候结束循环，可以从语句里面看出来

//func swapPairs(head *ListNode) *ListNode {
//	dummyhead := new(ListNode)
//	dummyhead.Next = head
//
//	cur := dummyhead
//	for cur.Next != nil && cur.Next.Next != nil {
//		temp1 := cur.Next.Next.Next
//		temp2 := cur.Next
//		cur.Next = cur.Next.Next
//		cur.Next.Next = temp2
//		temp2.Next = temp1
//		cur = cur.Next.Next
//	}
//	return dummyhead.Next
//}

//// 删除链表的倒数第N个节点
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if head == nil {
//		return head
//	}
//	//要找到第N个节点的前一个节点
//	//双指针：fast指针先向前前进n+1步，然后fast和slow一起向后前进，直到fast指向空节点
//	dummyhead := new(ListNode)
//	dummyhead.Next = head
//	fast, slow := dummyhead, dummyhead
//	//fast前进n+1步
//	n = n + 1
//	for n != 0 && fast != nil {
//		fast = fast.Next
//		n--
//	}
//	//fast,slow同时前进
//	for fast != nil {
//		fast = fast.Next
//		slow = slow.Next
//	}
//	//删除倒数第N个节点
//	slow.Next = slow.Next.Next
//	return head
//}

//func isAnagram(s string, t string) bool {
//	nums := make([]int, 26)
//	tt := []byte(t)
//	ss := []byte(s)
//
//	for _, v := range ss {
//		nums[v-'a']++
//	}
//	for _, m := range tt {
//		nums[m-'a']--
//	}
//	for i := 0; i < 26; i++ {
//		if nums[i] < 0 {
//			return false
//		}
//	}
//	return true
//}

//func intersection(nums1 []int, nums2 []int) []int {
//	mp := make(map[int]bool)
//	res := []int{}
//	for _, n2 := range nums2 {
//		mp[n2] = true
//	}
//	for _, n1 := range nums1 {
//		if mp[n1] {
//			res = append(res, n1)
//			mp[n1] = false
//		}
//	}
//	return res
////}
//
//func step(n int) int {
//	sum := 0
//	for n != 0 {
//		sum += (n % 10) * (n % 10)
//		n = n / 10
//	}
//	return sum
//}
//func isHappy(n int) bool {
//	mp := make(map[int]bool)
//	//一个set表
//	for n != 1 && !mp[n] {
//		mp[n] = true
//		n = step(n)
//	}
//	return n == 1
//}

//func twoSum(nums []int, target int) []int {
//	//要返回数组下标，所以要value存下标
//	//遍历过程中边寻找边添加
//	mp := make(map[int]int)
//
//	for index, num := range nums {
//		if preindex, ok := mp[target-num]; ok {
//			return []int{preindex, index}
//		} else {
//			mp[num] = index
//		}
//
//	}
//
//	return []int{}
//}

//// 出现的次数不能用bool，必须要用int
//func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
//
//	mp := make(map[int]int)
//	//构造哈希表
//	for i := 0; i < len(nums1); i++ {
//		for j := 0; j < len(nums2); j++ {
//			mp[nums1[i]+nums2[j]]++
//		}
//	}
//	res := 0
//
//	for i := 0; i < len(nums3); i++ {
//		for j := 0; j < len(nums4); j++ {
//			if mp[0-nums3[i]-nums4[j]] != 0 {
//				res += mp[0-nums3[i]-nums4[j]]
//			}
//		}
//	}
//	return res
//}

//func canConstruct(ransomNote string, magazine string) bool {
//	//因为范围为26个小写字母，所以使用数组
//	//将magazine生成一个哈希表
//	mp := make([]int, 26)
//
//	for _, m := range magazine {
//		mp[m-'a']++
//	}
//	for _, r := range ransomNote {
//		mp[r-'a']--
//		if mp[r-'a'] < 0 {
//			return false
//		}
//	}
//	return true
//
//}

// 三数之和
// 不推荐使用哈希表，去重麻烦
// 使用双指针
//func threeSum(nums []int) [][]int {
//	sort.Ints(nums)
//	res := [][]int{}
//
//	for i := 0; i < len(nums)-2; i++ {
//		//剪枝
//		if nums[i] > 0 {
//			break
//		}
//		//去重
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		//双指针
//		left := i + 1
//		right := len(nums) - 1
//		for left < right {
//			sum := nums[i] + nums[left] + nums[right]
//			if sum > 0 {
//				right--
//			} else if sum < 0 {
//				left++
//			} else {
//				res = append(res, []int{nums[i], nums[left], nums[right]})
//				for left < right && nums[left] == nums[left+1] {
//					left++
//				}
//				for left < right && nums[right] == nums[right-1] {
//					right--
//				}
//				left++
//				right--
//			}
//		}
//
//	}
//	return res
//}

// 四数之和 双指针法：可以使时间复杂度降1级，因为双指针遍历一次，相当于做了2层循环的事情
//func fourSum(nums []int, target int) [][]int {
//	sort.Ints(nums)
//	res := [][]int{}
//
//	for k := 0; k < len(nums)-3; k++ {
//		//剪枝
//		if nums[k] > target && target >= 0 {
//			break
//		}
//		//去重
//		if k > 0 && nums[k] == nums[k-1] {
//			continue
//		}
//		for i := k + 1; i < len(nums)-2; i++ {
//			//剪枝
//			if nums[i]+nums[k] > target && target > 0 {
//				break
//			}
//			//去重
//			if i > k+1 && nums[i] == nums[i-1] {
//				continue
//			}
//			//双指针
//			left := i + 1
//			right := len(nums) - 1
//			for left < right {
//				sum := nums[k] + nums[i] + nums[left] + nums[right]
//				if sum > target {
//					right--
//				} else if sum < target {
//					left++
//				} else {
//					res = append(res, []int{nums[k], nums[i], nums[left], nums[right]})
//					for left < right && nums[left] == nums[left+1] {
//						left++
//					}
//					for left < right && nums[right] == nums[right-1] {
//						right--
//					}
//					left++
//					right--
//				}
//			}
//		}
//
//	}
//	return res
//}

//}
//
//// 要找到翻转的起始位置
//func reverseStr(s string, k int) string {
//	ss := []byte(s)
//	for i := 0; i < len(ss); i += 2 * k {
//		if len(s)-i >= k {
//			reverseString(ss[i : i+k])
//		} else {
//			reverseString(ss[i:len(ss)])
//		}
//	}
//	return string(ss)
//}

//func replaceNumber(ss []byte) string {
//
//	oldsize := len(ss)
//	//统计数字个数
//	count := 0
//	for _, v := range ss {
//		if v >= '0' && v <= '9' {
//			count++
//		}
//	}
//	//扩充切片长度
//	tempstr := []byte("number")
//	for count > 0 {
//		ss = append(ss, []byte("     ")...)
//	}
//
//	//双指针：从后往前进行赋值
//	left := oldsize - 1
//	right := len(ss) - 1
//	for left <= right {
//		rightshift := 1
//		if ss[left] >= '0' && ss[right] <= '9' {
//			for i, str := range tempstr {
//				ss[right-len(tempstr)+1+i] = str
//			}
//			rightshift = len(tempstr)
//		} else {
//			ss[right] = ss[left]
//		}
//		//更新指针
//		right -= rightshift
//		left -= 1
//	}
//	return string(ss)
//}
//
//func main() {
//	var ss []byte
//	fmt.Scanln(&ss)
//
//	newString := replaceNumber(ss)
//
//	fmt.Println(newString)
//
//}

//}
//
//func reverseWords(s string) string {
//	ss := []byte(s)
//	//去除字符串中的空格（相当于删除数组中的元素）
//	slow := 0
//	for fast := 0; fast < len(ss); fast++ {
//		if ss[fast] != ' ' {
//			if slow != 0 {
//				ss[slow] = ' '
//				slow++
//			}
//			for fast < len(s) && ss[fast] != ' ' {
//				ss[slow] = ss[fast]
//				//更新slow
//				slow++
//				fast++
//			}
//		}
//	}
//	//更新字符串长度
//	ss = ss[0:slow]
//	//翻转整个字符串
//	reverseString(ss)
//	//翻转每个单词(切片)，要找到起始位置，空格位置
//	last := 0
//	for i := 0; i < len(ss); i++ {
//		if ss[i] == ' ' {
//			reverseString(ss[last:i])
//			//更新last
//			last = i + 1
//		}
//	}
//	reverseString(ss[last:len(ss)])
//	return string(ss)
//}

//func reverseString(s []byte) {
//	left := 0
//	right := len(s) - 1
//
//	for left < right {
//		s[left], s[right] = s[right], s[left]
//		left++
//		right--
//	}
//}
//
//// 右旋字符串
//func right1(s string, k int) string {
//	ss := []byte(s)
//	reverseString(ss)
//	reverseString(ss[:k])
//	reverseString(ss[k:])
//
//	return string(ss)
//}

// KMP
// 生成next数组，回退
//func getNext(next []int, s string) {
//	//双指针
//	//初始化
//	j := 0
//	next[0] = 0
//
//	for i := 1; i < len(s); i++ {
//		//前缀和后缀不等
//		for j > 0 && s[i] != s[j] {
//			j = next[j-1]
//		}
//		//前缀和后缀相等
//		if s[i] == s[j] {
//			j++
//		}
//		//更新next数组
//		next[i] = j
//	}
//}
//
//// 使用next数组跳过不必要的比较，回退
//func strStr(haystack string, needle string) int {
//	n := len(needle)
//	next := make([]int, n)
//	getNext(next, needle)
//
//	j := 0
//
//	for i := 0; i < len(haystack); i++ {
//		for j > 0 && haystack[i] != needle[j] {
//			j = next[j-1]
//		}
//		if haystack[i] == needle[j] {
//			j++
//		}
//		if j == n {
//			return i - j + 1
//		}
//	}
//	return -1
//
//}

////用栈实现队列
////要使用2个栈，用来改变顺序
//
//type MyQueue struct {
//	stackIn  *MyStack
//	stackOut *MyStack
//}
//
//func Constructor() MyQueue {
//	return MyQueue{
//		stackIn:  &MyStack{},
//		stackOut: &MyStack{},
//	}
//}
//
//func (this *MyQueue) Push(x int) {
//	this.stackIn.Push(x)
//}
//
//func (this *MyQueue) Pop() int {
//	//首先检查stackOut是否为空
//	if this.stackOut.Empty() {
//		//将stackIn中元素弹出并假如到stackOut中
//		for !this.stackIn.Empty() {
//			val := this.stackIn.Pop()
//			this.stackOut.Push(val) //把stackin中元素假如stackOut中
//		}
//	}
//	return this.stackOut.Pop()
//}
//
//func (this *MyQueue) Peek() int {
//	res := this.stackOut.Pop()
//	this.stackOut.Push(res)
//	return res
//}
//
//func (this *MyQueue) Empty() bool {
//	return this.stackIn.Empty() && this.stackOut.Empty()
//}

//// 通过切片实现一个栈
//type MyStack []int
//
//func (s *MyStack) Push(v int) {
//	*s = append(*s, v)
//}
//
//func (s *MyStack) Pop() int {
//	val := (*s)[len(*s)-1]
//	*s = (*s)[:len(*s)-1]
//	return val
//}
//
//func (s *MyStack) Peek() int {
//	return (*s)[len(*s)-1]
//}
//
//func (s *MyStack) Size() int {
//	return len(*s)
//}
//
//func (s *MyStack) Empty() bool {
//	return s.Size() == 0
//}

//func removeDuplicates(s string) string {
//	//创建一个栈，遍历字符串，若当前字符和栈顶元素不相等，那么入栈，
//	//如果相等则弹出栈顶元素
//	var stack []byte //栈
//	for i := 0; i < len(s); i++ {
//		//栈为空 或者 与栈顶（字符数组的尾部）不相等
//		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
//			stack = append(stack, s[i])
//		} else {
//			//如果相等，弹出栈顶元素（删除字符数组尾部元素）
//			stack = stack[:len(stack)-1]
//		}
//	}
//	return string(stack)
//}

//func evalRPN(tokens []string) int {
//	//给我一个后缀表达式然后求值
//	//消除，定义一个栈，遇到数字就入栈，遇到运算符，就弹出栈顶的两个元素进行运算，并将运算结果加入到栈中
//	var stack []int
//
//	for i := 0; i < len(tokens); i++ {
//		if tokens[i] == "+" {
//			n1 := stack[len(stack)-1]
//			n2 := stack[len(stack)-2]
//			stack = stack[:len(stack)-2]
//			stack = append(stack, n1+n2)
//		} else if tokens[i] == "-" {
//			n1 := stack[len(stack)-1]
//			n2 := stack[len(stack)-2]
//			stack = stack[:len(stack)-2] //弹出栈顶的两个元素
//			stack = append(stack, n2-n1)
//		} else if tokens[i] == "*" {
//			n1 := stack[len(stack)-1]
//			n2 := stack[len(stack)-2]
//			stack = stack[:len(stack)-2] //弹出栈顶的两个元素
//			stack = append(stack, n2*n1)
//		} else if tokens[i] == "/" {
//			n1 := stack[len(stack)-1]
//			n2 := stack[len(stack)-2]
//			stack = stack[:len(stack)-2] //弹出栈顶的两个元素
//			stack = append(stack, n2/n1)
//		} else {
//			sb, _ := strconv.Atoi(tokens[i])
//			stack = append(stack, sb)
//		}
//	}
//
//	return stack[len(stack)-1]
//}

// 用队列来实现栈

//type MyStack struct {
//	queue []int
//}
//
//func Constructor() MyStack {
//	return MyStack{
//		queue: make([]int, 0),
//	}
//}
//
//func (this *MyStack) Push(x int) {
//	this.queue = append(this.queue, x)
//}
//
//func (this *MyStack) Pop() int {
//	size := len(this.queue)
//	size--
//	for size != 0 {
//		//如何弹出队头元素，因为是切片，所以可以随意截取，和选取元素
//		val := this.queue[0]
//		this.queue = this.queue[1:]
//		this.queue = append(this.queue, val)
//		size--
//	}
//	//弹出队头元素
//	res := this.queue[0]
//	this.queue = this.queue[1:]
//	return res
//}
//
//func (this *MyStack) Top() int {
//	//n := len(this.queue) - 1
//	//
//	//for n != 0 {
//	//	val := this.queue[0]
//	//	this.queue = this.queue[1:]
//	//	this.queue = append(this.queue, val)
//	//	n--
//	//}
//	//res := this.queue[0]
//	//this.queue = this.queue[1:]
//	//this.queue = append(this.queue, res) //!!!!!!!!!!访问完后要记得放回去（恢复到初始状态），因为下次访问时，还要把前size-1个元素添加到队列后面
//	//return res
//	res := this.Pop()
//	this.queue = append(this.queue, res)
//	return res
//}
//
//func (this *MyStack) Empty() bool {
//	return len(this.queue) == 0
//}

//func isValid(s string) bool {
//	//括号匹配就用栈，遇到左括号就往栈里放右括号，遇到右括号就和栈顶元素匹配
//	stack := []byte{} //定义一个栈
//	ss := []byte(s)
//	for i := 0; i < len(ss); i++ {
//		if ss[i] == '(' {
//			stack = append(stack, ')')
//		} else if ss[i] == '[' {
//			stack = append(stack, ']')
//		} else if ss[i] == '{' {
//			stack = append(stack, '}')
//		} else { //ss[i]为右括号
//			//如果右括号多余，或者右括号与栈顶元素不匹配
//			if len(stack) == 0 || stack[len(stack)-1] != ss[i] {
//				return false
//			} else {
//				stack = stack[:len(stack)-1] //!!!!!!!!!!!!!!要记得出栈
//			}
//		}
//	}
//	//遍历结束，如果左括号多余,栈不为空，切片长度不为0
//	return len(stack) == 0
//}
//单调队列

//type MyQueue struct {
//	queue []int
//}
//
//func NewMyQueue() *MyQueue {
//	return &MyQueue{
//		queue: make([]int, 0),
//	}
//}
//
//func (m *MyQueue) Front() int {
//	return m.queue[0]
//}
//
//func (m *MyQueue) Back() int {
//	return m.queue[len(m.queue)-1]
//}
//
//func (m *MyQueue) Empty() bool {
//	return len(m.queue) == 0
//}
//func (m *MyQueue) Pop(val int) {
//	//要弹出的val和队头元素相等，那么弹出
//	if !m.Empty() && val == m.Front() {
//		m.queue = m.queue[1:]
//	}
//}
//
//func (m *MyQueue) Push(val int) {
//	//如果队尾元素小于val，那么弹出，直到队尾元素大于等于val，然后添加val
//	for !m.Empty() && m.Back() < val {
//		m.queue = m.queue[:len(m.queue)-1]
//	}
//	m.queue = append(m.queue, val)
//}
//
//func (m *MyQueue) getMaxValue() int {
//	return m.Front()
//}
//
//func maxSlidingWindow(nums []int, k int) []int {
//	que := NewMyQueue()
//	res := []int{}
//
//	//将前k个元素放入队列中
//	for i := 0; i < k; i++ {
//		que.Push(nums[i])
//	}
//	//记录前k个元素的最大值
//	res = append(res, que.getMaxValue())
//
//	for i := k; i < len(nums); i++ {
//		//移除滑动窗口最前面的元素
//		que.Pop(nums[i-k])
//		que.Push(nums[i])
//		res = append(res, que.getMaxValue())
//	}
//	return res
//}
//
//// 构建小顶堆
//// 因为是键值对，所以要构建一个2维切片，行数不固定，列数固定为2
//type IntHeap [][2]int
//
//// 实现最小堆独有的方法
//// 长度
//func (h IntHeap) Len() int {
//	return len(h)
//}
//
//// 比大小
//func (h IntHeap) Less(i, j int) bool {
//	return h[i][1] < h[j][1] //比较出现的频率
//}
//
//// 交换位置
//func (h IntHeap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//// 类型的push
//func (h *IntHeap) Push(x interface{}) {
//	*h = append((*h), x.([2]int)) //[2]int长度为2的数组
//}
//
//// 类型的pop
//func (h *IntHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	res := old[n-1]
//	*h = old[:n-1]
//	return res
//}
//
//// 前k个高频元素
//// 首先，要统计数字的频率（使用哈希表）
//func topKFrequent(nums []int, k int) []int {
//	mp := make(map[int]int)
//
//	//统计元素频率
//	for _, num := range nums {
//		mp[num]++
//	}
//
//	h := &IntHeap{}
//	heap.Init(h)
//	//哈希表中键值对入堆
//	for key, value := range mp {
//		heap.Push(h, [2]int{key, value})
//		if h.Len() > k {
//			heap.Pop(h)
//		}
//	}
//
//	res := make([]int, k)
//	//按顺序返回堆中元素
//	for i := 0; i < k; i++ {
//		res[k-i-1] = heap.Pop(h).([2]int)[0]
//	}
//	//要对哈希表中的键值对按值进行排序
//	//使用小顶堆
//	return res
//}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//func Traversal(cur *TreeNode, res *[]int) { //确定传入参数
//	//确定返回条件
//	if cur == nil {
//		return
//	}
//	//确定递归逻辑
//	*res = append((*res), cur.Val)
//	Traversal(cur.Left, res)
//	Traversal(cur.Right, res)
//}
//func preorderTraversal(root *TreeNode) []int {
//	//确定传入参数和返回值
//	res := &[]int{}
//	Traversal(root, res)
//	return *res
//}

//func preorderTraversal(root *TreeNode) (res []int) {
//	var traversal func(node *TreeNode)
//	traversal = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		res = append(res, node.Val)
//		traversal(node.Left)
//		traversal(node.Right)
//	}
//	traversal(root)
//	return res
//}
//
//func postorderTraversal(root *TreeNode) (res []int) {
//	var traversal func(node *TreeNode)
//	traversal = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		traversal(node.Left)
//		traversal(node.Right)
//		res = append(res, node.Val)
//	}
//	traversal(root)
//	return res
//}
//
//func inorderTraversal(root *TreeNode) (res []int) {
//	var traversal func(node *TreeNode) //
//	traversal = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		traversal(node.Left)
//		res = append(res, node.Val)
//		traversal(node.Right)
//	}
//	traversal(root)
//	return res
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//// 二叉树前序遍历-迭代法（使用栈）
//func preorderTraversal(root *TreeNode) (res []int) {
//	//中左右
//	//栈中先进右节点
//
//	//定义一个栈，用来存放节点指针
//	st := []*TreeNode{}
//	//结果数组已经在返回值中定义好了
//
//	//遍历根节点，将根节点压入栈中
//	st = append(st, root)
//
//	for len(st) != 0 {
//		node := st[len(st)-1] //取得栈顶元素
//		st = st[:len(st)-1]   //弹出栈顶元素
//		//如果栈顶元素是空指针,则跳过，因为遇到叶子节点它们的左右孩子节点也被压入栈中
//		if node == nil {
//			continue
//		}
//		res = append(res, node.Val) //处理该节点，将值加入结果数组（中）
//		st = append(st, node.Right) //先将右孩子节点压入栈中
//		st = append(st, node.Left)  //然后再将左节点压入栈中，这样出栈顺序才是(左右)
//	}
//	return res
//
//}

//// 二叉树后序遍历-迭代法（使用栈）
//func preorderTraversal(root *TreeNode) (res []int) {
//	//左右中
//	//栈中先进左节点 由前序的中左右 变成 中右左 最后翻转结果数组变成 左右中
//
//	//定义一个栈，用来存放节点指针
//	st := []*TreeNode{}
//	//结果数组已经在返回值中定义好了
//
//	//遍历根节点，将根节点压入栈中
//	st = append(st, root)
//
//	for len(st) != 0 {
//		node := st[len(st)-1] //取得栈顶元素
//		st = st[:len(st)-1]   //弹出栈顶元素
//		//如果栈顶元素是空指针,则跳过，因为遇到叶子节点它们的左右孩子节点也被压入栈中
//		if node == nil {
//			continue
//		}
//		res = append(res, node.Val) //处理该节点，将值加入结果数组（中）
//		//st = append(st, node.Right) //先将右孩子节点压入栈中
//		//st = append(st, node.Left)  //然后再将左节点压入栈中，这样出栈顺序才是(左右)
//		st = append(st, node.Left) //和前序顺序相反
//		st = append(st, node.Right)
//	}
//	return ReverseSlice(res)
//}
//
//func ReverseSlice(slice []int) []int {
//	left := 0
//	right := len(slice) - 1
//
//	for left < right {
//		slice[left], slice[right] = slice[right], slice[left]
//		left++
//		right--
//	}
//	return slice
//}

// // 中序遍历-迭代法（使用栈）
// // 左中右
//
//	func inorderTraversal(root *TreeNode) (res []int) {
//		//用栈来存放遍历过的节点
//
//		//定义一个栈
//		st := []*TreeNode{}
//
//		//定义一个指针用来进行遍历
//		cur := root
//
//		for cur != nil || len(st) != 0 { //指针不指向空 或者 栈不为空
//			if cur != nil {
//				st = append(st, cur) //将遍历过的节点压入栈中
//				cur = cur.Left       //然后一路向左（左）
//			} else { //如果指向空节点，那么我们弹出栈顶元素进行处理，将其值加入到结果数组中
//				cur = st[len(st)-1]        //将栈顶元素取出
//				st = st[:len(st)-1]        //弹出栈顶元素
//				res = append(res, cur.Val) //(中)
//				cur = cur.Right            //（右）
//			}
//		}
//		return res
//
// }

//func levelOrder(root *TreeNode) (res [][]int) {
//	//层序遍历使用队列
//
//	//定义一个队列
//	queue := []*TreeNode{}
//	if root == nil {
//		return res
//	}
//	//首先，将根节点加入队列
//	queue = append(queue, root)
//
//	for len(queue) > 0 {
//		length := len(queue)
//		tmpArr := []int{}
//		for i := 0; i < length; i++ {
//			//取队头元素
//			node := queue[0]
//			//弹出队头元素
//			queue = queue[1:]
//			//将该节点的值加入到结果数组
//			tmpArr = append(tmpArr, node.Val) //(中)二叉树的中部得到处理
//			//将当前节点的左右孩子(非空节点)入栈
//			if node.Left != nil {
//				queue = append(queue, node.Left) //（左）二叉树的左部得到处理
//
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right) // （右）二叉树的右部得到处理（一颗二叉树已经处理完，逻辑完整）
//			}
//		}
//		res = append(res, tmpArr)
//	}
//
//	return ReverseRes(res)
//}
//
////自底向上层序遍历
////思路：使用从上向下层序遍历，最后将结果数组逆置一下
//func ReverseRes(res [][]int) [][]int {
//	left := 0
//	right := len(res) - 1
//
//	for left < right {
//		res[left], res[right] = res[right], res[left]
//		left++
//		right--
//	}
//
//	return res
//}

//func rightSideView(root *TreeNode) (res []int) {
//	//层序遍历，然后只将每层的最后一个节点的值加入结果数组
//
//	//层序遍历
//	//使用队列
//	queue := []*TreeNode{}
//
//	if root == nil {
//		return res
//	}
//
//	//将根节点加入队列
//	queue = append(queue, root)
//
//	for len(queue) > 0 {
//		//记录每层的长度
//		length := len(queue)
//
//		for i := 0; i < length; i++ {
//			//取队头元素
//			node := queue[0]
//			//弹出队头元素
//			queue = queue[1:]
//			//如果左右孩子非空，则将其加入到队列中
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//			if i == length-1 { //每层的最后一个节点
//				//将该节点的值放入结果数组
//				res = append(res, node.Val)
//			}
//		}
//	}
//	return res
//}

//// 二叉树的层平均值
//func averageOfLevels(root *TreeNode) (res []float64) {
//	//层序遍历，求出每层的总和再除以每层的长度，得到平均值，然后加入到结果数组
//
//	//队列
//	queue := []*TreeNode{}
//
//	if root == nil {
//		return res
//	}
//
//	//将根节点加入队列
//	queue = append(queue, root)
//
//	for len(queue) > 0 {
//		//记录每层长度
//		length := len(queue)
//		sum := 0
//		for i := 0; i < length; i++ {
//			//取队头元素
//			node := queue[0]
//			//弹出队头元素
//			queue = queue[1:]
//
//			//如果左右孩子非空，则将其加入队列中
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//			//对该层求和
//			sum += node.Val
//		}
//		res = append(res, float64(sum)/float64(length))
//	}
//
//	return res
//}
//
//type Node struct {
//	Val      int
//	Children []*Node
//}
//
//func levelOrder(root *Node) (res [][]int) {
//	//层序遍历，求出每层的总和再除以每层的长度，得到平均值，然后加入到结果数组
//
//	//队列
//	queue := []*Node{}
//
//	if root == nil {
//		return res
//	}
//
//	//将根节点加入队列
//	queue = append(queue, root)
//
//	for len(queue) > 0 {
//		//记录每层长度
//		length := len(queue)
//		tmpArr := []int{}
//		for i := 0; i < length; i++ {
//			//取队头元素
//			node := queue[0]
//			//弹出队头元素
//			queue = queue[1:]
//
//			//遍历孩子数组，如果节点非空，将节点加入队列
//			for _, child := range node.Children {
//				if child != nil {
//					queue = append(queue, child)
//				}
//			}
//			//将该节点的值放入结果数组
//			tmpArr = append(tmpArr, node.Val)
//		}
//		res = append(res, tmpArr)
//	}
//
//	return res
//}

//// 求每层的最大值
//func largestValues(root *TreeNode) (res []int) {
//	//层序遍历，求出每层的总和再除以每层的长度，得到平均值，然后加入到结果数组
//
//	//队列
//	queue := []*TreeNode{}
//
//	if root == nil {
//		return res
//	}
//
//	//将根节点加入队列
//	queue = append(queue, root)
//
//	for len(queue) > 0 {
//		//记录每层长度
//		length := len(queue)
//		max := math.MinInt
//		for i := 0; i < length; i++ {
//			//取队头元素
//			node := queue[0]
//			//弹出队头元素
//			queue = queue[1:]
//
//			//如果左右孩子非空，则将其加入队列中
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//			//最大值
//			if node.Val > max {
//				max = node.Val
//			}
//		}
//		res = append(res, max)
//	}
//
//	return res
//}

//type Node struct {
//	Val      int
//	Children []*Node
//}
//
//func postorder(root *Node) (res []int) {
//	//匿名函数
//	var traversal func(node *Node)
//	traversal = func(node *Node) {
//		//确定返回条件
//		if node == nil {
//			return
//		}
//
//		//后序遍历 左右中
//
//		for _, child := range node.Children {
//			traversal(child)
//		} //左右
//		res = append(res, node.Val) //中
//
//		return
//	}
//	traversal(root)
//	return res
//}

////判断一颗二叉树是否对称
//func isSymmetric(root *TreeNode) bool {
//
//	//确定递归函数的形参，返回值（形参 左右子树指针，返回值 bool）
//	//使用匿名函数来实现递归函数
//	var traversal func(left *TreeNode, right *TreeNode) bool
//	traversal = func(left *TreeNode, right *TreeNode) bool { //要遍历两颗二叉树
//		//确定递归结束条件
//		if left == nil && right != nil {
//			return false
//		} else if left != nil && right == nil {
//			return false
//		} else if left == nil && right == nil {
//			return true
//		} else if left.Val != right.Val {
//			return false
//		}
//		//必须使用后序遍历（左右中）
//		outside := traversal(left.Left, right.Right)
//		inside := traversal(left.Right, right.Left)
//		res := outside && inside //（中）
//		return res
//	}
//
//	return traversal(root.Left, root.Right)
//
//}

//func maxDepth(root *TreeNode) (height int) {
//	//确定递归形参和返回值（形参：node 二叉树节点指针 返回值：树的高度）
//	var traversal func(node *TreeNode) int
//	traversal = func(node *TreeNode) int {
//		//确定递归返回条件
//		if node == nil {
//			return 0
//		}
//		left_height := traversal(node.Left)
//		right_height := traversal(node.Right)
//		height = max(left_height, right_height)
//
//		return height + 1
//	}
//	return traversal(root)
//
//}

//func minDepth(root *TreeNode) int {
//	//确定递归形参和返回值（形参：node 二叉树节点指针 返回值：树的高度）
//	var traversal func(node *TreeNode) int
//	traversal = func(node *TreeNode) int {
//		//确定递归返回条件
//		if node == nil {
//			return 0
//		}
//		left_height := traversal(node.Left)
//		right_height := traversal(node.Right)
//		if node.Left == nil && node.Right != nil {
//			return 1 + right_height
//		} else if node.Left != nil && node.Right == nil {
//			return 1 + left_height
//		} else {
//			height := 1 + int(min(left_height, right_height))
//			return height
//		}
//	}
//	return traversal(root)
//
//}

//func countNodes(root *TreeNode) int {
//	var traversal func(node *TreeNode) int //确定递归传入参数和返回值
//	traversal = func(node *TreeNode) int {
//		//终止条件（节点为空）
//		if node == nil {
//			return 0
//		}
//		//单层递归逻辑
//		left_num := traversal(node.Left)
//		right_num := traversal(node.Right)
//		return left_num + right_num + 1 //中
//	}
//	return traversal(root)
//}

//func isBalanced(root *TreeNode) bool {
//	var traversal func(node *TreeNode) (bool, int)
//	//形参和返回值
//	traversal = func(node *TreeNode) (bool, int) {
//		//终止条件
//		if node == nil {
//			return true, 0
//		}
//		//单层逻辑
//		isleftbalanced, left_height := traversal(node.Left)
//		isrightbalanced, right_height := traversal(node.Right)
//
//		if isrightbalanced && isleftbalanced && int(math.Abs(float64(left_height-right_height))) < 2 {
//			return true, 1 + int(max(left_height, right_height))
//		} else {
//			return false, -1
//		}
//
//	}
//	res, _ := traversal(root)
//	return res
//}

//func isBalanced(root *TreeNode) bool {
//	var traversal func(node *TreeNode) int
//	traversal = func(node *TreeNode) int {
//		if node == nil {
//			return 0
//		}
//
//		left_height := traversal(node.Left)
//		if left_height == -1 {
//			return -1
//		}
//		right_height := traversal(node.Right)
//		if right_height == -1 {
//			return -1
//		}
//		//中
//		res := 0
//		if int(math.Abs(float64(right_height-left_height))) > 1 {
//			res = -1
//		} else {
//			res = 1 + int(max(left_height, right_height))
//		}
//		return res
//	}
//	result := traversal(root)
//	if result == -1 {
//		return false
//	} else {
//		return true
//	}
//}

//func findBottomLeftValue(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	//层序遍历 使用队列
//	queue := []*TreeNode{}
//	res := [][]int{}
//	//根节点入队
//	queue = append(queue, root)
//
//	for len(queue) > 0 {
//		length := len(queue)
//		tmpArr := []int{}
//		for i := 0; i < length; i++ {
//			//取队头元素
//			node := queue[0]
//			//弹出队头元素
//			queue = queue[1:]
//			//将节点值加入结果数组
//			tmpArr = append(tmpArr, node.Val)
//			//如果左右孩子非空，则加入队列
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//		}
//		res = append(res, tmpArr)
//	}
//
//	return res[len(res)-1][0]
//}

//func binaryTreePaths(root *TreeNode) []string {
//	//不需要处理左右子树的返回值，所以没有返回值
//	var traversal func(node *TreeNode, path *[]int, res *[]string) //path用来存放单条路径，res用来存放所有路径
//	traversal = func(node *TreeNode, path *[]int, res *[]string) {
//		*path = append(*path, node.Val) //要先将叶子节点的值加入到路径数组中
//		//终止条件
//		if node.Left == nil && node.Right == nil {
//			paths := ""
//			for i := 0; i < len(*path)-1; i++ {
//				paths += strconv.Itoa((*path)[i])
//				paths += "->"
//			}
//			paths += strconv.Itoa((*path)[len(*path)-1])
//			*res = append(*res, paths)
//			return
//		}
//		//左右
//		if node.Left != nil {
//			traversal(node.Left, path, res)
//			*path = (*path)[:len(*path)-1] //回溯
//		}
//
//		if node.Right != nil {
//			traversal(node.Right, path, res)
//			*path = (*path)[:len(*path)-1] //回溯
//		}
//	}
//	path := []int{}
//	res := []string{}
//	traversal(root, &path, &res)
//	return res
//}

//func binaryTreePaths(root *TreeNode) []string {
//	res := make([]string, 0)
//	var travel func(node *TreeNode, s string) //s用来存单条路径
//	travel = func(node *TreeNode, s string) {
//		if node.Left == nil && node.Right == nil {
//			v := s + strconv.Itoa(node.Val)
//			res = append(res, v)
//			return
//		}
//		s = s + strconv.Itoa(node.Val) + "->"
//		if node.Left != nil {
//			travel(node.Left, s)
//			//s是字符串，是不可变的，每次递归调用都会生成一个新的局部拷贝，所以保证了返回时的回溯，每次递归返回时都恢复到了，递归调用前s 的状态
//			//体现出回溯
//		}
//		if node.Right != nil {
//func hasPathSum(root *TreeNode, targetSum int) bool {
//	if root == nil {
//		return false
//	}
//
//	return travel(root, targetSum)
//
//}
//
//func travel(node *TreeNode, count int) bool {
//	if node.Left == nil && node.Right == nil && count == 0 {
//		return true
//	}
//	if node.Left == nil && node.Right == nil && count != 0 {
//		return false
//	}
//
//	if node.Left != nil {
//		if travel(node.Left, count-node.Val) {
//			return true
//		}
//	}
//	if node.Right != nil {
//		if travel(node.Right, count-node.Val) {
//			return true //表示如果是true就一直向上返回
//		}
//	}
//	return false
//}
//			travel(node.Right, s)
//		}
//	}
//	travel(root, "")
//	return res
//
