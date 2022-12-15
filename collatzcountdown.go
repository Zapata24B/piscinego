package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	} else {
		i := 1
		for {
			if start%2 == 0 {
				start /= 2
			} else {
				start = (3*start + 1) / 2
			}
			i++
			if start == 1 {
				i++
				break
			}
		}
		return i
	}
}
