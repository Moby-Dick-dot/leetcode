package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	//树形dp
	//dp长度为2，表示对于当前节点有两个状态，0表示不偷当前节点，1表示偷当前节点
	//每一层递归，系统栈都会保存每个节点的状态dp
	//dp[0]表示不偷当前节点，能过获得的最高金额
	//dp[1]表示偷当前节点，能获得的最高金额
	dp := robtree(root)
	return max(dp[0], dp[1])
}

func robtree(cur *TreeNode) (dp []int) {
	if cur == nil {
		return []int{0, 0}
	}

	leftdp := robtree(cur.Left)
	rightdp := robtree(cur.Right)
	//偷当前节点所能获取的最大金额
	val1 := cur.Val + leftdp[0] + rightdp[0] //当前节点值加上不偷左、右孩子节点所能获得的最大金额
	//不偷当前节点所能获取的最大金额
	val2 := max(leftdp[0], leftdp[1]) + max(rightdp[0], rightdp[1]) //左孩子节点既可以偷也可以不偷，所以求出最大值即可

	return []int{val2, val1}
}
