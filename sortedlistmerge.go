package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	head := n1
	for current := n2; current != nil; {
		next := current.Next
		head = SortListInsertNode(n1, current)
		current = next
	}
	return head
}
