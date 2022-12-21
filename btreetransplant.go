package piscine

func BTreeTransplant(root, node, newNode *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	oldNode := node
	if node.Parent == nil {
		root = newNode
	} else if node == node.Parent.Left {
		oldNode.Parent.Left = newNode
	} else {
		oldNode.Parent.Right = newNode
	}
	newNode.Parent = node.Parent
	return root
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
