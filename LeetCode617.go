package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	return mergeTrees1(root1, root2)
}

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 != nil && root2 == nil {
		return root1
	}
	if root1 == nil && root2 != nil {
		return root2
	}
	node := &TreeNode{Val: 0, Left: nil, Right: nil}
	node.Val = root1.Val + root2.Val

	node.Left = mergeTrees1(root1.Left, root2.Left)
	node.Right = mergeTrees1(root1.Right, root2.Right)

	return node
}
