package piscine

func ListForEachIf(l *List, f func(*NodeL), condition func(*NodeL) bool) {
	a := l.Head
	for a != nil {
		if condition(a) {
			f(a)
		}
		a = a.Next
	}
}

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
