package piscine

import "fmt"

func ChunkCoulibaly(slice []int, size int) {
	chunk := [][]int(nil)
	tmp := make([]int, size)
	rest := len(slice) % size
	j := 0
	for i := 0; i < len(slice); i++ {
		tmp[j] = slice[i]
		j++
		if j == size {
			chunk = append(chunk, tmp)
			tmp = make([]int, size)
			j = 0
		}
	}
	if rest != 0 {
		chunk = append(chunk, slice[len(slice)-rest:])
	}
	fmt.Println(chunk)
}
