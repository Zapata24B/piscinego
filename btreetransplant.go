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
	oldNode.Parent = node.Parent
	return root
}
