package main

func convertBST(root *TreeNode) *TreeNode {
	var traversal func(node *TreeNode) *TreeNode
	sum := 0

	traversal = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Right = traversal(node.Right)

		node.Val += sum
		sum = node.Val
		node.Left = traversal(node.Left)

		return node
	}
	return traversal(root)
}
