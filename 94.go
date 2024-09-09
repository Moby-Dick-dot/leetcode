package main

// 中序遍历：左中右
func inorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}

	var track func(node *TreeNode)
	track = func(node *TreeNode) {
		if node == nil {
			return
		}

		track(node.Left)
		res = append(res, node.Val)
		track(node.Right)
	}
	track(root)

	return res
}
