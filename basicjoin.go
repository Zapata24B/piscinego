package piscine

func BasicJoin(elems []string) string {
	str := ""
	for i := 0; i < len(elems); i++ {
		str = Concat(str, elems[i])
	}
	return str
}
