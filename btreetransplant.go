package piscine

func BTreeTransplant(root, oldNode, newNode *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root = BTreeDeleteNode(root, oldNode)
	root = BTreeInsertNode(root, newNode)
	return root
	// if root.Data == oldNode.Data {
	// 	return root
	// } else if root.Data > oldNode.Data {
	// 	return BTreeTransplant(root.Left, oldNode, newNode)
	// } else {
	// 	return BTreeTransplant(root.Right, oldNode, newNode)
	// }
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		} else {
			temp := BTreeMin(root.Right)

			root.Data = temp.Data
			root.Right = BTreeDeleteNode(root.Right, temp)
		}
	}
	return root
}

func BTreeInsertNode(root, newNode *TreeNode) *TreeNode {
	if root == nil {
		return newNode
	}
	if newNode.Data < root.Data {
		root.Left = BTreeInsertNode(root.Left, newNode)
		root.Left.Parent = root
	} else if newNode.Data > root.Data {
		root.Right = BTreeInsertNode(root.Right, newNode)
		root.Right.Parent = root
	}
	return root
}
