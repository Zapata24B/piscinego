package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref, Next: nil}

	for current, prev := l, l; current != nil; prev, current = current, current.Next {
		if current.Data > data_ref {
			newNode.Next = current
			if prev == current {
				return newNode
			} else {
				prev.Next = newNode
				break
			}
		} else if current.Next == nil {
			current.Next = newNode
			break
		}
	}
	return l
}
