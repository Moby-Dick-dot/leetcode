package main

func wordBreak(s string, wordDict []string) bool {
	//构造一个哈希表
	wordDictset := make(map[string]bool)

	for _, w := range wordDict {
		wordDictset[w] = true
	}

	//dp[i]表示字符串s的前i个字符能否拆成worddict
	dp := make([]bool, len(s)+1)
	//初始化
	dp[0] = true

	for i := 1; i <= len(s); i++ { //遍历背包容量
		for j := 0; j < i; j++ { //遍历物品（根据排列取得结果）
			if dp[j] == true && wordDictset[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}
