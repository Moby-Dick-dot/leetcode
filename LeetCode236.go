package main

//
//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	return lowestCommonAncestor1(root, p, q)
//
//}
//
//func lowestCommonAncestor1(node *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
//	if node == nil {
//		return nil
//	}
//	//根据左右子树如果都不为空，那么这个节点就是p和q 的公共祖先
//	if node == p || node == q {
//		return node
//	}
//
//	left := lowestCommonAncestor1(node.Left, p, q)
//	right := lowestCommonAncestor1(node.Right, p, q)
//
//	if left != nil && right != nil {
//		return node
//	} else if left != nil && right == nil {
//		return left
//	} else if left == nil && right != nil {
//		return right
//	}
//	return nil
//
//}
