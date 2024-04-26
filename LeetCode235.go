package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor1(root, p, q)
}

func lowestCommonAncestor1(node *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	//if (node.Val >= p.Val && node.Val <= q.Val) || (node.Val >= q.Val && node.Val <= p.Val) {
	//	return node
	//}

	if node.Val > p.Val && node.Val > p.Val {
		left := lowestCommonAncestor1(node.Left, p, q)
		if left != nil {
			return left
		}
	}

	if node.Val < p.Val && node.Val < q.Val {
		right := lowestCommonAncestor1(node.Right, p, q)
		if right != nil {
			return right
		}
	}

	return node
}
